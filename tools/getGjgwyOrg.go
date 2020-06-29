package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
	"io/ioutil"
	"regexp"
	"strings"

	"net/http"
	"os"
)

const (
	pageUrl = `http://www.gjgwy.org/mrlx/%d.html` //列表地址
	lineKey = "--------------------------------"
)

var (
	pageNo    = 1   //列表页面
	limitPage = 384 //尾页
	rowNo     = 1
)

//http://www.gjgwy.org/202006/449355.html
func main() {
	//准备数据库
	databases.InitMysqlDB()
	//1.获取所有题目的地址,panic终止
	//for{
	//	mrlx()
	//	pageNo++
	//	if pageNo > limitPage{
	//		break
	//	}
	//}
	//2.把所有连接的内容，通过class拆分，第一个jumbotron_well_gj 是题目，第二个jumbotron_well_gj 是解析
	//f, err := os.Open("listUrl.txt")
	//defer f.Close()
	//if err != nil {
	//	panic(err)
	//}
	//input := bufio.NewScanner(f)
	//for input.Scan() {
	//	line := input.Text()
	//	fmt.Println("line:", line)
	//	jumbotronWell(rowNo, line)
	//	rowNo++
	//}
	//3.题目内容处理
	//filepath.Walk("c/", func(path string, info os.FileInfo, err error) error {
	//	if info.IsDir() {
	//		return nil
	//	}
	//	if strings.Index(path, ".txt") == -1 {
	//		return nil
	//	}
	//	tm(path)
	//	return nil
	//})
	tm("c/row_10000.txt")
}

//列表页面获取连接
func mrlx() {
	listUrl := fmt.Sprintf(pageUrl, pageNo)
	fmt.Println("listUrl:", listUrl)
	//get 请求
	resp, err := http.Get(listUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(resp)
	}
	//goquery处理内容
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	//准备写入文件
	f, err := os.OpenFile("listUrl.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()
	doc.Find(".main_left .main_con .list > li").Each(func(i int, selection *goquery.Selection) {
		if val, exists := selection.Find(".cnt > a").Attr("href"); exists {
			_, err := f.WriteString(fmt.Sprintln(val))
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}

//拆解内容
func jumbotronWell(rowNum int, wellUrl string) {
	//get 请求
	resp, err := http.Get(wellUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	//goquery处理内容
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	//准备写入文件
	saveFile := fmt.Sprintf("c/row_%d.txt", rowNum)
	f, _ := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()
	doc.Find(".main_content .jumbotron_well_gj").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			//第一个是题目内容
			selection.Find(".panel-body").Each(func(j int, s *goquery.Selection) {
				ret, _ := s.Html()
				f.WriteString(fmt.Sprintln(trimHtml(ret)))
				f.WriteString(fmt.Sprintln(lineKey))
			})
		} else {
			//第二个是解析内容
			selection.Find(".question-analysis").Each(func(j int, s *goquery.Selection) {
				f.WriteString(fmt.Sprintln(s.Text()))
				f.WriteString(fmt.Sprintln(lineKey))
			})
		}
	})
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

//题目的处理
func tm(rowFile string) {
	//读取文件内容
	fileContent, err := ioutil.ReadFile(rowFile)
	if err != nil {
		return
	}
	//如果内容包含"http"不要，可能包含图片
	if bytes.Contains(fileContent, []byte("http")) {
		return
	}
	//如果内容包含"&gt"不要，不定是个什么东西
	if bytes.Contains(fileContent, []byte("&gt")) {
		return
	}
	//如果内容包含"F."不要，我们最多支持5个选项
	if bytes.Contains(fileContent, []byte("F.")) {
		return
	}
	//根据分隔符拆分
	fileContentList := bytes.Split(fileContent, []byte(lineKey))
	//如果长度是奇数，不要

	if (len(fileContentList)-1)%2 != 0 {
		return
	}
	//计数器，eg. 总长度是10，第一题的解析在第6行
	gLen := (len(fileContentList) - 1) / 2
	//模型容器
	var mList []*models.ChoiceQuestion
	for i := 0; i < gLen; i++ {
		//题干处理
		//fmt.Println("问题：")
		err, tg, o1, o2, o3, o4, o5 := tgO(fileContentList[i])
		if err != nil {
			continue
		}
		//至少4个选项都有值
		if len(o1) == 0 || len(o2) == 0 || len(o3) == 0 || len(o4) == 0 {
			continue
		}
		//fmt.Println(string(tg))
		//fmt.Println(string(o1))
		//fmt.Println(string(o2))
		//fmt.Println(string(o3))
		//fmt.Println(string(o4))
		//fmt.Println(string(o5))
		//答案解析处理
		//fmt.Println("答案：", string(fileContentList[i+gLen]))
		err, daStr, daIndex, jx := dajx(fileContentList[i+gLen])
		if err != nil {
			continue
		}
		//fmt.Println("daStr:", string(daStr))
		//fmt.Println("daIndex:", daIndex)
		//fmt.Println("jx:", string(jx))

		//拼装模型
		m := new(models.ChoiceQuestion)
		m.Stem = string(tg)
		m.Score = 1
		m.Answer = string(daStr)
		m.Analysis = string(jx)
		m.AnswerIndex = uint(daIndex)
		m.Options = append(m.Options, &models.ChoiceOption{
			OptionType: models.StringOption,
			Item:       string(o1),
		}, &models.ChoiceOption{
			OptionType: models.StringOption,
			Item:       string(o2),
		}, &models.ChoiceOption{
			OptionType: models.StringOption,
			Item:       string(o3),
		}, &models.ChoiceOption{
			OptionType: models.StringOption,
			Item:       string(o4),
		})
		//如果第五个选项有值
		if len(o5) > 0 {
			m.Options = append(m.Options, &models.ChoiceOption{
				OptionType: models.StringOption,
				Item:       string(o5),
			})
		}
		mList = append(mList, m)
	}
	if len(mList) > 0 {
		jsonFileName := strings.ReplaceAll(rowFile, "c/", "cjson/")
		jsonFileName = strings.ReplaceAll(jsonFileName, ".txt", ".json")
		jsonContent, err := json.Marshal(mList)
		if err != nil {
			return
		}
		ioutil.WriteFile(jsonFileName, jsonContent, 0777)
		//写库
		//for _, m := range mList{
		//	err := databases.NewDB().Save(m).Error
		//	if err != nil{
		//		fmt.Println(err)
		//	}
		//}
	}
}

//题干
func tgO(content []byte) (err error, tg, o1, o2, o3, o4, o5 []byte) {
	//计算A.的位置,截取作为题干处理
	aIndex := bytes.Index(content, []byte("A."))
	if aIndex == -1 {
		err = fmt.Errorf("aIndex error")
		return
	}
	//干掉换行
	tg = deleteThNum(bytes.ReplaceAll(content[:aIndex], []byte("\n"), []byte("")))
	//计算B.的位置,截取作为A选项
	bIndex := bytes.Index(content, []byte("B."))
	if bIndex == -1 {
		err = fmt.Errorf("bIndex error")
		return
	}
	o1 = deleteLastOneLineCode(content[aIndex+2 : bIndex])
	//计算C.的位置,截取作为B选项
	cIndex := bytes.Index(content, []byte("C."))
	if cIndex == -1 {
		err = fmt.Errorf("cIndex error")
		return
	}
	o2 = deleteLastOneLineCode(content[bIndex+2 : cIndex])
	//计算D.的位置,截取作为C选项
	dIndex := bytes.Index(content, []byte("D."))
	if dIndex == -1 {
		err = fmt.Errorf("dIndex error")
		return
	}
	o3 = deleteLastOneLineCode(content[cIndex+2 : dIndex])
	//计算E.的位置,截取作为D选项,如果没有，直接到结尾算D选项
	eIndex := bytes.Index(content, []byte("E."))
	if eIndex == -1 {
		o4 = deleteLastOneLineCode(content[dIndex+2:])
	} else {
		o4 = deleteLastOneLineCode(content[dIndex+2 : eIndex])
		o5 = deleteLastOneLineCode(content[eIndex+2:])
	}
	return
}

//答案解析
func dajx(content []byte) (err error, daStr []byte, daIndex int, jx []byte) {
	content = deleteThNum(content)
	//冒号都变成英文，这样能减少解析位置的问题
	content = bytes.ReplaceAll(content, []byte("："), []byte(":"))
	//解析处理
	jxLen := bytes.Index(content, []byte("解析"))
	if jxLen == -1 {
		err = fmt.Errorf("jxLen")
		return
	}
	jx = bytes.TrimSpace(content[jxLen+len("解析: "):])
	//答案处理
	daLen := bytes.Index(content, []byte("答案:"))
	if daLen == -1 {
		err = fmt.Errorf("daLen")
		return
	}
	daStr = bytes.TrimSpace(content[daLen+len("答案:") : jxLen])
	daI := bytes.Index([]byte("ABCDE"), daStr)
	if daI == -1 {
		err = fmt.Errorf("daI")
		return
	}
	daIndex = daI
	return
}

//去掉最后一位换行符
func deleteLastOneLineCode(o []byte) []byte {
	if len(o) < 1 {
		return o
	}
	if o[len(o)-1] == '\n' {
		return o[:len(o)-1]
	}
	return o
}

//去掉题号1. 2. 3. 4. 5. 6. 7. 8. 9.
func deleteThNum(o []byte) []byte {
	if len(o) < 1 {
		return o
	}
	thList := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, th := range thList {
		if o[0] == th && o[1] == '.' {
			return o[2:]
		}
	}
	return o
}

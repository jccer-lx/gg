package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/impl"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/golayui/weight"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

type AdminTableData struct {
	ApiUrl    string            `json:"api_url"`
	Id        string            `json:"id"`
	ViewName  string            `json:"view_name"`
	FieldList []*AdminFieldData `json:"field_list"`
}

type AdminFieldData struct {
	Field string `json:"field"`
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

var adminTableMap map[string]impl.AdminModelImpl

func init() {
	adminTableMap = make(map[string]impl.AdminModelImpl)
	adminTableMap["demo"] = &models.Admin{}
	adminTableMap["demo1"] = &models.Admin{}
	adminTableMap["demo2"] = &models.Admin{}
}

func AdminIndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "layout/common.html", nil)
}

func AdminListView(c *gin.Context) {
	viewName := c.Param("view")
	logrus.Debug("AdminListView")
	adminTableData := new(AdminTableData)
	adminTableData.ViewName = viewName
	//map中取出对应结构
	fl, err := adminTableMap[viewName].GetTableFields()
	if err != nil {
		ErrorView(c)
		return
	}
	//字段
	for _, f := range fl {
		adminTableData.FieldList = append(adminTableData.FieldList, &AdminFieldData{
			Field: f.JsonTitle,
			Title: f.Title,
			Sort:  "true",
		})
	}
	//接口地址
	adminTableData.ApiUrl = fmt.Sprintf("/admin/api/%s", viewName)
	adminTableData.Id = fmt.Sprintf("%X", time.Now().UnixNano())
	c.HTML(http.StatusOK, "admin/list.html", adminTableData)
}

func AdminAddView(c *gin.Context) {
	viewName := c.Param("view")
	logrus.Debug("AdminAddView")
	m := adminTableMap[viewName]
	mHtml, err := addTpl(viewName, m)
	if err != nil {
		fmt.Println(err)
		ErrorView(c)
		return
	}
	fmt.Println("mHtml:", mHtml)
	//c.HTML(http.StatusOK, "admin/form.html", nil)
	c.String(http.StatusOK, mHtml)
}

func AdminEditView(c *gin.Context) {
	viewName := c.Param("view")
	logrus.Debug("AdminEditView")
	m := adminTableMap[viewName]
	mHtml, err := addTpl(viewName, m)
	if err != nil {
		ErrorView(c)
		return
	}
	c.String(http.StatusOK, mHtml, nil)
}

//add_form模板内容
func addTpl(viewName string, m impl.AdminModelImpl) (string, error) {
	savePath := "views/tmp"
	mHtmlPath := fmt.Sprintf("%s/%s.html", savePath, viewName)
	logrus.Debug("mHtmlPath:", mHtmlPath)
	//判断文件是否存在
	if ok, _ := helper.PathExists(mHtmlPath); ok {
		//读取内容
		logrus.Debug("读取内容")
		mHtml, err := ioutil.ReadFile(mHtmlPath)
		if err != nil {
			logrus.Error("读取内容error", err)
		}
		return string(mHtml), nil
	} else {
		logrus.Debug("文件不存在")
		//生成文件
		formLabelWeight, err := templateFile(m)
		if err != nil {
			logrus.Error("addTpl 生成文件错误：", err)
			return "", err
		}
		mHtml, err := formLabelWeight.Output()
		if err != nil {
			logrus.Error("addTpl 生成mHtml错误：", err)
			return "", err
		}
		//渲染元素
		formTpl, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.html", "views/admin", "add_form"))
		if err != nil {
			logrus.Error("addTpl 读取渲染内容错误：", err)
			return "", err
		}
		tmpl, err := template.New("").Parse(string(formTpl))
		if err != nil {
			logrus.Error("addTpl 渲染内容错误：", err)
			return "", err
		}
		builder := new(bytes.Buffer)
		err = tmpl.Execute(builder, map[string]string{
			"FormHtml": mHtml,
			"FormId":   formLabelWeight.Id,
		})
		if err != nil {
			logrus.Error("addTpl 执行渲染错误：", err)
			return "", err
		}
		logrus.Debug("strBuilder.String:", builder.String())
		err = ioutil.WriteFile(mHtmlPath, builder.Bytes(), 0666)
		if err != nil {
			logrus.Error("addTpl 写文件错误：", err)
			return "", err
		}
		return builder.String(), nil
	}
}

//edit_form模板内容
func editTpl(viewName string, m impl.AdminModelImpl) (string, error) {
	savePath := "views/tmp"
	mHtmlPath := fmt.Sprintf("%s/%s.html", savePath, viewName)
	logrus.Debug("mHtmlPath:", mHtmlPath)
	//判断文件是否存在
	if ok, _ := helper.PathExists(mHtmlPath); ok {
		//读取内容
		logrus.Debug("读取内容")
		mHtml, err := ioutil.ReadFile(mHtmlPath)
		if err != nil {
			logrus.Error("读取内容error", err)
		}
		return string(mHtml), nil
	} else {
		logrus.Debug("文件不存在")
		//生成文件
		formLabelWeight, err := templateFile(m)
		if err != nil {
			logrus.Error("readTpl 生成文件错误：", err)
			return "", err
		}
		mHtml, err := formLabelWeight.Output()
		if err != nil {
			logrus.Error("readTpl 生成mHtml错误：", err)
			return "", err
		}
		//渲染元素
		formTpl, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.html", "views/admin", "add_form"))
		if err != nil {
			logrus.Error("readTpl 读取渲染内容错误：", err)
			return "", err
		}
		tmpl, err := template.New("").Parse(string(formTpl))
		if err != nil {
			logrus.Error("readTpl 渲染内容错误：", err)
			return "", err
		}
		builder := new(bytes.Buffer)
		err = tmpl.Execute(builder, map[string]string{
			"FormHtml": mHtml,
			"FormId":   formLabelWeight.Id,
		})
		if err != nil {
			logrus.Error("readTpl 执行渲染错误：", err)
			return "", err
		}
		logrus.Debug("strBuilder.String:", builder.String())
		err = ioutil.WriteFile(mHtmlPath, builder.Bytes(), 0666)
		if err != nil {
			logrus.Error("readTpl 写文件错误：", err)
			return "", err
		}
		return builder.String(), nil
	}
}

//生成form模板文件
func templateFile(m impl.AdminModelImpl) (*weight.FormLabelWeight, error) {
	fl, err := m.GetTableFields()
	if err != nil {
		return nil, err
	}
	formLabelWeight := new(weight.FormLabelWeight)
	formLabelWeight.Label = m.TableName()
	formLabelWeight.FormWeight.Id = fmt.Sprintf("form%d", time.Now().UnixNano())
	for _, f := range fl {
		fiw := new(weight.FormItemWeight)
		formLabelWeight.FormWeight.Children = append(formLabelWeight.FormWeight.Children, fiw)
		fiw.Label = f.Title
		//根据类型生成对应html内容
		var item weight.FormItemWeightImpl
		//id name
		attr := weight.Attr{
			Name: f.JsonTitle,
			Id:   f.JsonTitle,
		}
		switch f.Type {
		case weight.CheckBox:
			//TODO option
			item = &weight.CheckboxWeight{
				OptionList: nil,
			}

		case weight.Select:
			item = &weight.SelectWeight{
				Attr:       attr,
				OptionList: nil,
			}
			//TODO option
		case weight.InputPassword:
			item = &weight.InputPasswordWeight{
				Attr:        attr,
				Placeholder: f.Title,
			}
		case weight.Textarea:
			item = &weight.TextareaWeight{
				Attr:        attr,
				Placeholder: f.Title,
			}
		case weight.Radio:
			item = &weight.RadioWeight{
				OptionList: nil,
			}
			//TODO option
		default:
			item = &weight.InputTextWeight{
				Attr:        attr,
				Placeholder: f.Title,
				IsDate:      false,
			}
		}
		fiw.Item = item
	}
	//添加提价按钮
	formLabelWeight.FormWeight.Children = append(formLabelWeight.FormWeight.Children, &weight.ButtonWeight{
		Attr: weight.Attr{
			Id: fmt.Sprintf("btn%d", time.Now().UnixNano()),
		},
		Title: "提交",
	})
	return formLabelWeight, nil
}

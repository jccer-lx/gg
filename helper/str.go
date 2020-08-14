package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//string -> uint
func String2Uint(str string) uint {
	i, err := strconv.Atoi(str)
	if err != nil {
		logrus.Error("String2Uint error:", err)
		//panic(err)
		return 0
	}
	return uint(i)
}

//string -> int
func String2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logrus.Error("String2Int error:", err)
		//panic(err)
		return 0
	}
	return i
}

//string -> float64
func String2Float64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		logrus.Error("String2Float64 error:", err)
		//panic(err)
		return 0
	}
	return f
}

//string -> int64
func String2Int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logrus.Error("String2Int error:", err)
		//panic(err)
		return 0
	}
	return i
}

func JsonNumber2Uint(jn json.Number) uint {
	i, err := jn.Int64()
	if err != nil {
		logrus.Error("JsonNumber2Uint error:", err)
		//panic(err)
		return 0
	}
	return uint(i)
}

func JsonNumber2Int(jn json.Number) int {
	i, err := jn.Int64()
	if err != nil {
		logrus.Error("JsonNumber2Int error:", err)
		//panic(err)
		return 0
	}
	return int(i)
}

func JsonNumber2Int64(jn json.Number) int64 {
	i, err := jn.Int64()
	if err != nil {
		logrus.Error("JsonNumber2Int64 error:", err)
		//panic(err)
		return 0
	}
	return i
}

func JsonNumber2Float64(jn json.Number) float64 {
	i, err := jn.Float64()
	if err != nil {
		logrus.Error("JsonNumber2Float64 error:", err)
		//panic(err)
		return 0.00
	}
	return i
}

func Switch2Int(switchStr string) int {
	if switchStr == "on" {
		return 1
	} else {
		return 0
	}
}

//uuid
func UUidV4() string {
	v := uuid.NewV4()
	return v.String()
}

//数字变字母0-A
func Num2Letter(num int) string {
	var (
		str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	slice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if num > 26 { //数据大于26需要进行拆分
		for {
			k = num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			num = (num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, num)
				break
			}
		}
	} else {
		return slice[num]
	}

	for _, value := range temp {
		str = slice[value] + str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return str
}

//字母变数字A-0
func Letter2Num(letter string) int {
	//用来匹配的字符A-Z
	slice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, l := range slice {
		if l == letter {
			return i
		}
	}
	return 9999999
}

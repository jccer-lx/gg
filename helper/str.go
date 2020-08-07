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

package helper

import (
	"crypto/md5"
	"encoding/hex"
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
		return 0
	}
	return uint(i)
}

//uuid
func UUidV4() string {
	v := uuid.NewV4()
	return v.String()
}

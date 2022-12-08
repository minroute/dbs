package helper

import (
	"math/rand"
	"strings"
	"time"
)

// 获取随机字母+数字组合字符串
func RandStr(length int) string {
	if length < 1 {
		return ""
	}
	char := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

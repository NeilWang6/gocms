package utils

import (
	"fmt"
	"regexp"
)
import "crypto/md5"
import "math/rand"
import "time"
import "strconv"
import "strings"

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

/**
获取随机字符串
*/
func GetRandomString(length int, number bool) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	if number {
		str = "0123456789"
	}
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetTimeRange(class string) (timeRange string) {
	classNum := strings.TrimLeft(class, "class")
	classNumInt, _ := strconv.Atoi(classNum)
	classNumInt = classNumInt / 10
	if classNumInt > 0 && classNumInt <= 7 {
		timeRange = "6-8"
	} else if classNumInt > 7 && classNumInt <= 14 {
		timeRange = "8-10"
	} else if classNumInt > 14 && classNumInt <= 21 {
		timeRange = "10-12"
	} else if classNumInt > 21 && classNumInt <= 28 {
		timeRange = "12-14"
	} else if classNumInt > 28 && classNumInt <= 35 {
		timeRange = "15-17"
	} else if classNumInt > 35 && classNumInt <= 42 {
		timeRange = "17-19"
	} else if classNumInt > 42 && classNumInt <= 49 {
		timeRange = "19-21"
	} else if classNumInt > 49 && classNumInt <= 56 {
		timeRange = "21-23"
	} else {
		timeRange = "0-0"
	}
	return
}

func GetDateByClass(class string) (date string) {
	classNum := strings.TrimLeft(class, "class")
	classNumInt, _ := strconv.Atoi(classNum)
	classNumInt = classNumInt / 10
	if classNumInt%7 == 0 {
		date = GetNextNDate(time.Now(), 7)
	} else {
		date = GetNextNDate(time.Now(), classNumInt%7)
	}
	return
}

func GetLengthNameByClass(class string) string {
	lengName := class + "Length"
	return lengName
}

func TrimHtml(src string) string {
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
	//去除所有的空格
	re, _ = regexp.Compile("&nbsp;")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}

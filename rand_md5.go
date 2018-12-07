package common

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

/**
* 获取随机字符串
* @length  int  随机数的长度
* @param  string   随机数的类型
* 类型说明：number:数字   small:小写字母   big:大写字母   word:大小写字母   smallnumber:小写字母和数字   bignumber:大写字母和数字   all:大小写字母和数字
* return	string		返回一个字符串
 */
func GetRadomString(length int, param string) string {
	str := ""
	result := ""
	switch param {
	case "number":
		str = "0123456789"
	case "small":
		str = "abcdefghijklmnopqrstuvwxyz"
	case "big":
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "word":
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "smallnumber":
		str = "abcdefghijklmnopqrstuvwxyz0123456789"
	case "bignumber":
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case "all":
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	}
	strLen := len(str)
	var star int
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		star = ran.Intn(strLen)
		result = result + Substr(str, star, 1)
	}
	return result
}

/**
* 返回指定长度的随机数
 */
func GetRand(l int) string {
	var res string
	var list = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	lens := len(list) - 1

	randKey := rand.Intn(lens)
	//干扰随机数规律
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for j := 0; j < r.Intn(9); j++ {
		randKey = rand.Intn(lens)
	}
	//正式生成随机数
	for i := 0; i < l; i++ {
		randKey = rand.Intn(lens)
		res = res + list[randKey]
	}
	return res
}

/**
* 生成两个数之间的随机数
* @min   int64  最小值
* @max   int64  最大值
* return  int64	 返回一个int64整形数字
 */
func RandomMaxAndMin(min, max int64) int64 {
	if min >= max {
		return max
	}
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := ran.Int63n(max-min) + min
	return res
}

/**
* 返回md5加密后的数据
 */
func GetMd5(source string) string {
	if len(source) < 1 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

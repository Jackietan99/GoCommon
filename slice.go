package common

import (
	"reflect"
)

/**
* 判断一个字符串在切片中是否存在
* @param		src			切片
* @param		value		数值
* return		bool			存在=true
 */
func SliceContains(src []string, value string) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

/*
循环赋值一个map,避免删除key的时候2个map一起删除
*/
func GetNewMap(mm map[string]string) map[string]string {
	newMm := map[string]string{}
	for k, v := range mm {
		newMm[k] = v
	}
	return newMm
}

/**
* 去除重复
 */
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

// Function to check if the the string given is in the array
func IsInArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

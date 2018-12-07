package common

import (
	"sort"
	"strconv"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringSlice) Less(i, j int) bool {
	one, _ := strconv.Atoi(s[i])
	two, _ := strconv.Atoi(s[j])
	if one < two {
		return true
	} else {
		return false
	}

}

/**
数字数组排序
*/
func SliceSort(aSource []string) []string {
	var aTarget StringSlice
	aTarget = append(aTarget, aSource...)

	sort.Sort(aTarget)

	return aTarget

}

//按某个字段排序
type sortByPayType []map[string]string

func (s sortByPayType) Len() int           { return len(s) }
func (s sortByPayType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByPayType) Less(i, j int) bool { return s[i]["pay_type"] < s[j]["pay_type"] }

//切片分组
func SplitSlice(list []map[string]string) [][]map[string]string {
	sort.Sort(sortByPayType(list))
	var returnData [][]map[string]string
	var tempArr1, tempArr2, tempArr3, tempArr4, tempArr5, tempArr6, tempArr7, tempArr8 []map[string]string

	for _, v := range list {
		switch v["pay_type"] {
		case "0":
			tempArr1 = append(tempArr1, v)
		case "1":
			tempArr2 = append(tempArr2, v)
		case "2":
			tempArr3 = append(tempArr3, v)
		case "3":
			tempArr4 = append(tempArr4, v)
		case "4":
			tempArr5 = append(tempArr5, v)
		case "5":
			tempArr6 = append(tempArr6, v)
		case "6":
			tempArr7 = append(tempArr7, v)
		case "7":
			tempArr8 = append(tempArr8, v)
		}
	}
	if len(tempArr1) > 0 {
		returnData = append(returnData, tempArr1)
	}
	if len(tempArr2) > 0 {
		returnData = append(returnData, tempArr2)
	}
	if len(tempArr3) > 0 {
		returnData = append(returnData, tempArr3)
	}
	if len(tempArr4) > 0 {
		returnData = append(returnData, tempArr4)
	}
	if len(tempArr5) > 0 {
		returnData = append(returnData, tempArr5)
	}
	if len(tempArr6) > 0 {
		returnData = append(returnData, tempArr6)
	}
	if len(tempArr7) > 0 {
		returnData = append(returnData, tempArr7)
	}
	if len(tempArr8) > 0 {
		returnData = append(returnData, tempArr8)
	}

	return returnData
}

//按某个字段排序
type sortByDisplayPayType []map[string]string

func (s sortByDisplayPayType) Len() int      { return len(s) }
func (s sortByDisplayPayType) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortByDisplayPayType) Less(i, j int) bool {
	return s[i]["display_pay_type"] < s[j]["display_pay_type"]
}

//切片分组
func SplitSliceUserFor(list []map[string]string) [][]map[string]string {
	sort.Sort(sortByDisplayPayType(list))

	returnData := make([][]map[string]string, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && list[i]["display_pay_type"] == list[j]["display_pay_type"]; j++ {
		}

		returnData = append(returnData, list[i:j])
		i = j
	}
	return returnData

	return returnData
}

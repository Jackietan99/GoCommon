package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
 * 拆分字符串成一个个字符,并且去重
 */
func DivideStrDuplicateRemoval(sBaseWinningNumber string) []string {
	var aStr []string
	mStr := map[string]int{}
	iLen := len(sBaseWinningNumber)
	for i := 0; i < iLen; i++ {

		//去重处理
		sEffective := sBaseWinningNumber[i : i+1]
		if _, ok := mStr[sEffective]; ok != true {
			mStr[sEffective] = i
			aStr = append(aStr, sEffective)
		} else {
			continue
		}

	}
	return aStr
}

/*
 * 拆分字符串成一个个字符,不去重
 */
func DivideStrNoDuplicateRemoval(sBaseWinningNumber string) []string {
	return strings.Split(sBaseWinningNumber, "")
}

/*
 * 获取两个数组交集
 */
func ArrayIntersect(aNumbers, aOther []string) []string {
	var aData []string
	for _, sNumber := range aNumbers {
		if InArray(sNumber, aOther) {
			aData = append(aData, sNumber)
		}
	}

	return aData
}

/**
 *获取两个数组差集
 */
func ArrayDiff(aNumbers, aOther []string) []string {
	var aData []string
	for _, sNumber := range aNumbers {
		if !InArray(sNumber, aOther) {
			aData = append(aData, sNumber)
		}
	}

	return aData
}

/*
 * 判断是否在数组中(strings)
 */
func InArray(sNumber string, aNumbers []string) bool {
	var bIsIn bool = false
	for _, sValue := range aNumbers {
		if sNumber == sValue {
			bIsIn = true
		}
	}

	return bIsIn
}

/*
 * 判断是否在数组中(int)
 */
func InArrayInt(iNumber int, aNumbers []int) bool {
	var bIsIn bool = false
	for _, iValue := range aNumbers {
		if iNumber == iValue {
			bIsIn = true
		}
	}

	return bIsIn
}

/*
 * 数组去重
 */
func ArrayUnique(aNumbers []string) []string {
	var aData []string

	if len(aNumbers) < 1024 {
		for i := range aNumbers {
			flag := true
			for j := range aData {
				if aNumbers[i] == aData[j] {
					flag = false // 存在重复元素，标识为false
					break
				}
			}
			if flag {
				// 标识为false，不添加进结果
				aData = append(aData, aNumbers[i])
			}
		}
	} else {
		tempMap := map[string]byte{} // 存放不重复主键
		for _, e := range aNumbers {
			l := len(tempMap)
			tempMap[e] = 0
			if len(tempMap) != l {
				// 加入map后，map长度变化，则元素不重复
				aData = append(aData, e)
			}
		}
	}

	return aData
}

/*
 *　统计数组中值的个数
 */
func ArrayCountValues(aNumbers []string) map[string]int {
	mStr := make(map[string]int)
	for _, sNumber := range aNumbers {
		if _, ok := mStr[sNumber]; ok == true {
			mStr[sNumber] += 1
		} else {
			mStr[sNumber] = 1
		}

	}

	return mStr
}

/*
 * 获取数组和值
 */
func ArraySum(aNumbers []string) int {
	var sum int = 0
	for _, sNumber := range aNumbers {
		iNumber, _ := strconv.Atoi(sNumber)
		sum += iNumber
	}
	return sum
}

/*
 * 获取数组值最大值
 */
func ArrayMax(aNumbers []string) int {
	var max int = 0
	for _, sNumber := range aNumbers {
		iNumber, _ := strconv.Atoi(sNumber)
		if iNumber > max {
			max = iNumber
		}
	}
	return max
}

/*
 * 获取数组值最大值(浮点数)
 */
func ArrayMaxF(aNumbers []string) float64 {
	var max float64 = 0
	fmt.Println("aNumbers==>", aNumbers)
	for _, sNumber := range aNumbers {
		fmt.Println("sNumber==>", sNumber)
		fNumber, _ := strconv.ParseFloat(sNumber, 64)
		fmt.Println("fNumber==>", fNumber)
		if fNumber > max {
			max = fNumber
		}
	}
	return max
}

/*
 * 获取数组值最小值
 */
func ArrayMin(aNumbers []string) int {
	min := aNumbers[0]
	iMin, _ := strconv.Atoi(min)
	for _, sNumber := range aNumbers {
		iNumber, _ := strconv.Atoi(sNumber)
		if iNumber < iMin {
			iMin = iNumber
		}
	}
	return iMin
}

/*
 * 获取数组值最小值,float64
 */
func ArrayMinF(aNumbers []string) float64 {
	min := aNumbers[0]
	fMin, _ := strconv.ParseFloat(min, 64)
	for _, sNumber := range aNumbers {
		fNumber, _ := strconv.ParseFloat(sNumber, 64)
		if fNumber < fMin {
			fMin = fNumber
		}
	}
	return fMin
}

/*
 * 获取map值最大值
 */
func MapMax(aNumbers map[string]int) int {
	var max int = 0
	for _, iNumber := range aNumbers {
		if iNumber > max {
			max = iNumber
		}
	}
	return max
}

/*
 * 获取map值最大值,string
 */
func MapMaxStr(aNumbers map[string]string) int {
	var max int = 0
	for _, sNumber := range aNumbers {
		iNumber, _ := strconv.Atoi(sNumber)
		if iNumber > max {
			max = iNumber
		}
	}
	return max
}

/*
 * 获取map值SUM,float64
 */
func MapSumF(mData map[string]string) float64 {
	var sum float64 = 0
	for _, sValue := range mData {
		fValue, _ := strconv.ParseFloat(sValue, 64)
		sum += fValue
	}
	return sum
}

/*
 * 获取map值SUM,int
 */
func MapSum(mData map[string]string) int {
	var sum int = 0
	for _, sValue := range mData {
		iValue, _ := strconv.Atoi(sValue)
		sum += iValue
	}
	return sum
}

/*
 * 获取map值最小值
 */
func MapMin(aNumbers map[string]int) int {
	var min int = 9999999

	for _, iNumber := range aNumbers {
		if iNumber < min {
			min = iNumber
		}
	}
	return min
}

/*
 * 获取map值最小值,float64
 */
func MapMinF(aNumbers map[string]float64) float64 {
	var min float64 = 9999999
	for _, fNumber := range aNumbers {
		if fNumber < min {
			min = fNumber
		}
	}
	return min
}

/*
 * 获取map值最大值,float64
 */
func MapMaxF(aNumbers map[string]float64) float64 {
	var max float64 = 0
	for _, fNumber := range aNumbers {
		if fNumber > max {
			max = fNumber
		}
	}
	return max
}

/*
 * 获取map值keys
 */
func MapKeys(aNumbers map[string]int, aSearchValue []string) []string {
	var arr []string
	for index, iValue := range aNumbers {
		if len(aSearchValue) > 0 {
			sValue := strconv.Itoa(iValue)
			if InArray(sValue, aSearchValue) {
				arr = append(arr, index)
			}
		} else {
			arr = append(arr, index)
		}
	}

	arr = ArrayUnique(arr)
	return arr
}

/*
 * 获取map值keys
 */
func MapKeysStr(aNumbers map[string]string) []string {
	var arr []string
	for index, _ := range aNumbers {
		arr = append(arr, index)
	}
	return arr
}

/*
 * 返回数组值keys，如果指定了值，只返回指定只的key
 *@param aNumbers []string 原数组
 *@param aSearchValue []string 指定值
 *
 */
func ArrayKeys(aNumbers []string, aSearchValue []string) []string {
	var arr []string
	for index, sValue := range aNumbers {
		sIndex := strconv.Itoa(index)
		if len(aSearchValue) > 0 {
			if InArray(sValue, aSearchValue) {
				arr = append(arr, sIndex)
			}
		} else {
			arr = append(arr, sIndex)
		}

	}

	arr = ArrayUnique(arr)
	return arr
}

/**
数组中检索某个值对应的key
*/
func ArraySearch(sNumber string, aNumbers []string) interface{} {
	for iKey, sItem := range aNumbers {
		iNumber, err := strconv.Atoi(sNumber)
		iItem, err := strconv.Atoi(sItem)
		if err != nil {
			return false
		}
		if iItem == iNumber {
			return iKey
		}

	}
	return false

}

/**
随机打乱数组的顺序
*/
func ArrayShuffle(aSource []string) []string {
	iSource := len(aSource)
	aTarget := make([]string, iSource)
	perm := rand.Perm(iSource)
	for i, v := range perm {
		aTarget[v] = aSource[i]
	}

	return aTarget
}

/*
 * 整形切片转换成字符串切片
 */
func IntArrTOStrArr(aInt []int) []string {
	aData := []string{}
	for _, iNumber := range aInt {
		aData = append(aData, fmt.Sprintf("%d", iNumber))
	}
	return aData
}

/*
 * 字符串切片转换成整形切片
 */
func StrArrTOIntArr(aStr []string) []int {
	aData := []int{}
	for _, sNumber := range aStr {
		iNumber, _ := strconv.Atoi(sNumber)
		aData = append(aData, iNumber)
	}
	return aData
}

/*
 *  用某个值从某位置开始填充数组到指定长度
 */
func ArrayFill(startIndex int, num int, value string) []string {
	m := make([]string, num)
	for i := 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

/*
 *  php array_column
 */
func ArrayColumn(mData []map[string]interface{}, sField string) []string {
	var aResult []string
	for _, m := range mData {
		for sKey, sValue := range m {
			if sKey == sField {
				aResult = append(aResult, sValue.(string))
			}
		}
	}
	return aResult
}

/*
 *  php array_fill
 */
func ArrayFillInt(startIndex int, num int, value int) []int {
	m := make([]int, num)
	for i := startIndex; i < num; i++ {
		m[startIndex] = value
	}
	return m
}

/*
 * php array_sum
 */
func IntArraySum(aNumbers []int) int {
	var sum int = 0
	for _, sNumber := range aNumbers {
		sum += sNumber
	}
	return sum
}

/*
 * php array_fill
 * 上面的写法有错误
 */
func ArrayFillIntTrue(startIndex int, num int, value int) map[int]int {
	m := map[int]int{}
	for i := 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

/*
 * 判断值在数组中出现的次数(strings)
 */
func InArrayNumStr(sNumber string, aNumbers []string) int {
	bIsIn := 0
	for _, sValue := range aNumbers {
		if sNumber == sValue {
			bIsIn++
		}
	}

	return bIsIn
}

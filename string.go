package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"common/uniqid"
)

/*
* 取指定长度的字符串，开始位置为0,如果从取2014010中取010的话Substr(4,3)
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/**
* 插入字符串
 */
func InsertStr(str, inser string, start int) string {
	res := str
	strLen := len(str)
	if start > 0 {
		if start == 0 {
			res = inser + res
		} else if strLen > start {
			pre := Substr(str, 0, start)
			suf := Substr(str, start, strLen-start)
			res = pre + inser + suf
		} else {
			res = res + inser
		}
	}

	return res
}

/**
* %s =  字符串
%q = 输出带"的字符串
%p = 输出指针
%d  = 整形
%b  = 二进制
%x  = 16进制
%t  = 布尔
%f  = 浮点   %.2f = 2位小数点
*/
func InterfaceToString(inter interface{}) string {
	res := ""
	switch inter := inter.(type) {
	case bool:
		res = fmt.Sprintf("%t", inter)
	case int:
		res = fmt.Sprintf("%d", inter)
	case int64:
		res = fmt.Sprintf("%d", inter)
	case float64:
		res = strconv.FormatFloat(inter, 'f', -1, 64)
	case byte:
		res = fmt.Sprintf("%b", inter)
	case string:
		res = fmt.Sprintf("%s", inter)
	case *bool:
		res = fmt.Sprintf("%p", inter)
	case *int:
		res = fmt.Sprintf("%p", inter)
	case *int64:
		res = fmt.Sprintf("%p", inter)
	case *float64:
		res = fmt.Sprintf("%p", inter)
	case *string:
		res = fmt.Sprintf("%p", inter)
	default:
		res = ""
	}
	return res
}

func ByteToString(b []byte) string {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return string(b[0:i])
		}
	}
	return string(b)
}

func BankStr(str string) string {
	str_len := len(str)
	new_str := string(str[(str_len - 4):str_len])
	return "***************" + new_str
}

/*
 * 返回手机/邮箱等
 * 这种格式 135****1234
 */
func PrivacyInfoBasic(str string) string {

	new_str := ""
	str_len := len(str)

	salt := "****"
	salt_len := len(salt)

	if str_len > 2 {

		var str_c float64
		str_c = float64(str_len) / float64(salt_len)

		f_str_int_s := math.Ceil(str_c) //向上取整
		str_int_s := int(f_str_int_s)

		//f_str_int_x := math.Floor(str_c) //向下取整
		//str_int_x := int(f_str_int_x)

		str_prefix := string(str[0:str_int_s])
		str_suffix := string(str[(str_len - 1):str_len])
		if (salt_len + str_int_s) < str_len {
			str_suffix = string(str[(salt_len + str_int_s):str_len])
		}

		new_str = str_prefix + salt + str_suffix
	} else {

		new_str = salt
	}
	return new_str
}

/*
 * 隐私信息保护
 * 手机/邮箱/QQ
 */
func PrivacyInfo(str string) string {
	new_str := ""

	str_len := len(str)

	if str_len > 0 {

		is_email := strings.Contains(str, "@")
		if is_email {

			new_str_arr := strings.Split(str, "@")
			if len(new_str_arr) > 1 {

				new_str_prefix := PrivacyInfoBasic(new_str_arr[0])
				new_str_suffix := new_str_arr[1]

				new_str = new_str_prefix + "@" + new_str_suffix
			}

		} else {

			new_str = PrivacyInfoBasic(str)
		}
	}

	return new_str
}

/*
 获取一个值在总数中的百分比
*/
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}

/*
 右填充
*/
func RightPad(str string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = str + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

/*
 左填充
*/
func LeftPad(str string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + str
	return retStr[(len(retStr) - overallLen):]
}

/**
将字符串转换为大写开始的驼峰
*/
func CamelCase(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = Ucwords(str)
	return strings.Replace(str, " ", "", -1)
}

/**
  字符串首字母大写
*/
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

/**
字符串首字母小写
*/
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

/**
单词首字母大写
*/
func Ucwords(str string) string {
	return strings.Title(str)
}

// php number_format()
// decimals: Sets the number of decimal points.
// decPoint: Sets the separator for the decimal point.
// thousandsSep: Sets the thousands separator.
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j, _ := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

func Uniqid(prefix string, bMoreEntropy bool) string {

	return uniqid.New(uniqid.Params{prefix, bMoreEntropy})

}

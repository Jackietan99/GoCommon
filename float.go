package common

import (
	"fmt"
	"math"
	"strconv"
)

/*
float64 向下取整 小数点后2位
return float64
*/
func F64_low_2(f64 float64) float64 {
	f64_cheng_100 := f64 * 100
	f64_cheng_100_low := F64_low(f64_cheng_100)
	f64_chu_100 := f64_cheng_100_low / 100
	return f64_chu_100
}

/*
float64 向下取整 小数点后2位
return string
*/
func F64_low_2_str(f64 float64) string {
	f64_chu_100 := F64_low_2(f64)
	f64_to_str := fmt.Sprintf("%.2f", f64_chu_100)
	return f64_to_str
}

/*
float64 向下取整
return float64
*/
func F64_low(f64 float64) float64 {
	return math.Floor(f64)
}

/*
返回绝对值
*/
func AbsFloat(f64 float64) float64 {
	new_f64 := f64
	if f64 < 0 {
		new_f64 = 0 - f64
	}
	return new_f64
}

//字符串转浮点数
func Str2Float64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_FORMAT_YMDHIS string = "2006-01-02 15:04:05"
	DATE_FORMAT_YMD    string = "2006-01-02"
)

/**
* 判断两个字符串日期相差的天数
 */
func DifferDays(startDate, endDate string) int {
	//startDateStr := "2015-10-07 15:20:10"
	//endDate := time.Now().Format("2006-01-02 15:04:05")
	//将初始日期处理为0时0分0秒
	startT, _ := FormatDateCCT("2006-01-02 15:04:05", startDate)
	startStr := startT.Format("2006-01-02 00:00:00")

	//将日期字符串转成时间格式
	startStrT, _ := FormatDateCCT("2006-01-02 00:00:00", startStr)
	endT, _ := FormatDateCCT("2006-01-02 15:04:05", endDate)
	//将时间格式转成时间戳
	startUnix := startStrT.Unix()
	endUnix := endT.Unix()
	//通过Unix时间戳获取两个日期相差的天数
	//dayNum := (endUnix - startUnix) / 86400
	unixCount := int(endUnix - startUnix)
	var dayNum int

	dayNum = (unixCount) / 86400
	return dayNum
}

/**
* 统一日期格式
 */
func FormatDateString(str, strFormat, dateFormat string, start, length int) string {
	dateStr := Substr(str, start, length)
	d, _ := time.ParseDuration("+12h")
	//endDate := "2015-10-13 15:05:23"
	dateTime, _ := FormatDateCCT(strFormat, dateStr)
	if strings.Contains(str, "PM") {
		dateTime = dateTime.Add(d)
	}
	formatStr := dateTime.Format(dateFormat)
	return formatStr
}

/**
* 计算时差返回字符串
* 字符串类型:   2006-01-02 15:04:05
 */
func TimeDiffString(dateStr, timeDiff string) string {
	res := dateStr
	d, _ := time.ParseDuration(timeDiff)
	dateTime, err := FormatDateCCT("2006-01-02 15:04:05", dateStr)
	if err == nil {
		dateTime = dateTime.Add(d)
		res = dateTime.Format("2006-01-02 15:04:05")
	}
	return res
}

/*
 * 格式化字符串为时间戳
 * @date 待转化为时间戳的字符串
 * @format 转化所需模板 如 20060102
 */
func FormatStr2Unix(date string, format string) (int64, error) {
	loc, _ := time.LoadLocation("Local")                    //重要：获取时区
	theTime, err := time.ParseInLocation(format, date, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0, err
	}
	return theTime.Unix(), nil
}

/*
 * 时间戳转日期
 * @time_unix 时间戳
 * @format 转化所需模板 如 20060102
 */
func FormatUnix2Str(time_unix int64, format string) string {
	dataTimeStr := time.Unix(time_unix, 0).Format(format)
	return dataTimeStr
}

/*
 * 格式化字符串北京时间
 */
func FormatDateCCT(format, date string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")                    //重要：获取时区
	theTime, err := time.ParseInLocation(format, date, loc) //使用模板在对应时区转化为time.time类型
	return theTime, err
}

/*
 * 获取
 * 2016-11-xx 00:00:00
 * 2016-11-xx 23:59:59
 */
func GetSDateAndEDateSJC(sjc int) (string, string) {

	nowTime := time.Now()

	add_day := fmt.Sprintf("%dh", sjc)
	th_bet_dt, _ := time.ParseDuration(add_day)
	sTime := nowTime.Add(th_bet_dt)
	sDate := sTime.Format(DATE_FORMAT_YMD)
	sDate += " 00:00:00"

	eDate := nowTime.Format(DATE_FORMAT_YMD)
	eDate += " 23:59:59"

	return sDate, eDate
}

/*
 * 获取现在日期时间
 */
func GetNowDatetime(format string) string {
	return time.Now().Format(format)
}

/**
 * 时间戳转换为时间字符串
 * @param	sTimeStamp	string 时间戳

 * @param	dateFormat	string	时间格式 eg:2006-01-02 03:04:05 ,20060102 03:04:05, 20060102030405
 * @return 	formatStr	string	格式化后时间
 */
func DateFormat(sTimeStamp string, dateFormat string) string {
	timeInt64, err := strconv.ParseInt(sTimeStamp, 10, 64)
	formatStr := ""
	if err == nil {
		timeStamp := time.Unix(timeInt64, 0)
		formatStr = timeStamp.Format(dateFormat)
	}
	return formatStr
}

/**
* 时间字符串转换为时间戳(目前只支持 yyyy-mm-dd hh:ii:ss 和 yyyymmddhhiiss两种格式 )
*
* @param	dateStr	string 时间
* @return 	timeStamp	时间戳
 */
func Strtotime(dateStr string) interface{} {
	var timeStamp interface{}
	dateTime, err := time.Parse(DATE_FORMAT_YMDHIS, dateStr)
	if err == nil {
		timeStamp = dateTime.Unix()
	} else {
		dateTime, err := time.Parse("20060102150405", dateStr)
		if err != nil {
			timeStamp = false
		}
		timeStamp = dateTime.Unix()

	}

	return timeStamp
}

/*
* 改变日期
* @time_date 要改变的日期
* @format 转化所需模板 如 20060102
* @years 增加或减少的年
* @months 增加或减少的月
* @days 增加或减少的日
 */
func ChangeDate(time_date, format string, years int, months int, days int) string {
	time_type, _ := time.Parse(format, time_date)
	new_time_type := time_type.AddDate(years, months, days)
	new_time_date := new_time_type.Format(format)
	return new_time_date
}

/*
增加或者减少当前时间
@cha_str	= -3h;当前时间减少3小时
@format	返回时间的格式
*/
func ChangeDateByNow(cha_str, format string) string {
	dt, _ := time.ParseDuration(cha_str)
	return time.Now().Add(dt).Format(format)
}

/*
获取今天	时间戳整型
*/
func GetTodayUnix() (int64, int64) {

	today_ymd := GetNowDatetime(DATE_FORMAT_YMD)

	today_morning := today_ymd + " 00:00:00"
	today_evening := today_ymd + " 23:59:59"

	today_morning_time, _ := FormatStr2Unix(today_morning, DATE_FORMAT_YMDHIS)
	today_evening_time, _ := FormatStr2Unix(today_evening, DATE_FORMAT_YMDHIS)

	return today_morning_time, today_evening_time
}

/*
获取昨天	datetime
*/
func GetYesterdayDateTime() (string, string) {

	today_ymd := time.Now().AddDate(0, 0, -1).Format(DATE_FORMAT_YMD)

	today_morning := today_ymd + " 00:00:00"
	today_evening := today_ymd + " 23:59:59"

	return today_morning, today_evening
}

/*
获取今天	时间戳字符串
*/
func GetTodayUnixStr() (string, string) {
	today_morning_time, today_evening_time := GetTodayUnix()
	today_morning_time_str := InterfaceToString(today_morning_time)
	today_evening_time_str := InterfaceToString(today_evening_time)

	return today_morning_time_str, today_evening_time_str
}

/**
* 获取本周的开始和结束时间戳
* @return		本周一开始时间
* @return		本周日结束时间
 */
func GetWeekStimeEtimeUnixStr() (string, string) {
	mondayTimeStr := GetMonday()
	mondayTimeStart := mondayTimeStr + " 00:00:00"
	mondayTimeEnd := mondayTimeStr + " 23:59:59"
	mondayStartUnix, _ := FormatDateCCT("2006-01-02 15:04:05", mondayTimeStart)
	mondayEndUnix, _ := FormatDateCCT("2006-01-02 15:04:05", mondayTimeEnd)

	sundayEndUnix := mondayEndUnix.Unix() + 86400*6

	mondayStr := fmt.Sprintf("%d", mondayStartUnix.Unix())
	sundayStr := fmt.Sprintf("%d", sundayEndUnix)
	return mondayStr, sundayStr
}

/**
 *获取上周的开始和结束时间戳
 *
 * @return		上周一开始时间
 * @return		上周日结束时间
 */
func GetLastWeekStartAndEndTimeUnixStr() (string, string) {
	sCurrentMondayDate := GetMonday()

	oTime, _ := time.Parse(DATE_FORMAT_YMD, sCurrentMondayDate)

	//计算上周的开始
	lastWeekStart := oTime.AddDate(0, 0, -7)
	sLastWeekStart := lastWeekStart.Format(DATE_FORMAT_YMD)
	sLastWeekStart = sLastWeekStart + " 00:00:00"
	iLastWeekStart, _ := FormatStr2Unix(sLastWeekStart, DATE_FORMAT_YMDHIS)
	mondayStr := fmt.Sprintf("%d", iLastWeekStart)

	//计算过上周结束时间
	lastWeekEnd := oTime.AddDate(0, 0, -1)
	sLastWeekEnd := lastWeekEnd.Format(DATE_FORMAT_YMD)
	sLastWeekEnd = sLastWeekEnd + " 23:59:59"
	iLastWeekEnd, _ := FormatStr2Unix(sLastWeekEnd, DATE_FORMAT_YMDHIS)
	sundayStr := fmt.Sprintf("%d", iLastWeekEnd)

	return mondayStr, sundayStr
}

/*
*获取本周一的日期
*
*@return 2017-04-17
 */
func GetMonday() string {

	nowTime := time.Now()

	sCurrentDay := GetDayOfTheWeek()

	nowDay := nowTime.Format(DATE_FORMAT_YMD) //今天

	mondayDate := "" //返回的 周1的日期

	switch sCurrentDay {

	case "Sunday": //周7
		mondayDate = ChangeDateByNow("-144h", DATE_FORMAT_YMD)

	case "Monday": //周1
		mondayDate = nowDay

	case "Tuesday": //周2
		mondayDate = ChangeDateByNow("-24h", DATE_FORMAT_YMD)

	case "Wednesday": //周3
		mondayDate = ChangeDateByNow("-48h", DATE_FORMAT_YMD)

	case "Thursday": //周4
		mondayDate = ChangeDateByNow("-72h", DATE_FORMAT_YMD)

	case "Friday": //周5
		mondayDate = ChangeDateByNow("-96h", DATE_FORMAT_YMD)

	case "Saturday": //周6
		mondayDate = ChangeDateByNow("-120h", DATE_FORMAT_YMD)

	}

	return mondayDate
}

/**
 *现在星期几
 */
func GetDayOfTheWeek() string {
	currentTime := time.Now()
	day := currentTime.Weekday()
	return day.String()

}

/**
 * 获得选定月份的开始和结束时间戳
 * @param sStr string 如 "2018-06"
 * return string string
 */
func GetFormulateStimeEtimeUnixStr(sStr string) (string, string) {
	sDate := sStr + "-01"
	iDate, _ := time.Parse(DATE_FORMAT_YMD, sDate)
	year, month, _ := iDate.Date()
	tFormulate := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	thisMonthEnd := tFormulate.AddDate(0, +1, -1)
	startStr := fmt.Sprintf("%d", tFormulate.Unix())
	endStr := fmt.Sprintf("%d", thisMonthEnd.Unix())

	return startStr, endStr
}

/**
* 获得当月的开始和结束时间戳
 */
func GetMonthStimeEtimeUnixStr() (string, string) {
	year, month, _ := time.Now().Date()
	//得到本月的第一天
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	//向后加一个月
	nextMonth := thisMonth.AddDate(0, +1, 0)
	//后一个月向前推一天，得到本月最后一天
	thisMonthEnd := nextMonth.AddDate(0, 0, -1)

	eYear, eMonth, eDay := thisMonthEnd.Date()
	thisMonthEnd = time.Date(eYear, eMonth, eDay, 23, 59, 59, 0, time.Local)

	startStr := fmt.Sprintf("%d", thisMonth.Unix())
	endStr := fmt.Sprintf("%d", thisMonthEnd.Unix())
	return startStr, endStr
}

/**
* 获得下個月的开始和结束时间戳
 */
func GetNextMonthStimeEtimeUnixStr() (string, string) {

	year, month, _ := time.Now().Date()

	//得到本月的第一天
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	//向后加一个月
	nextMonth := thisMonth.AddDate(0, +1, 0)

	//后一个月向前推一天，得到本月最后一天
	thisMonthEnd := nextMonth.AddDate(0, +1, -1)

	eYear, eMonth, eDay := thisMonthEnd.Date()
	thisMonthEnd = time.Date(eYear, eMonth, eDay, 23, 59, 59, 0, time.Local)

	startStr := fmt.Sprintf("%d", nextMonth.Unix())
	endStr := fmt.Sprintf("%d", thisMonthEnd.Unix())
	return startStr, endStr
}

/*
格式化当前时间格式为其他格式

ChangeDate("2018-05-24 00:59:00", "2006-01-02 15:04:05", "2006-01-02")
return 2018-05-24
*/
func ChangeDate2NewFormat(date_time, format, to_format string) string {

	loc, _ := time.LoadLocation("Local")                       //重要：获取时区
	theTime, _ := time.ParseInLocation(format, date_time, loc) //使用模板在对应时区转化为time.time类型
	new_time_date := theTime.Format(to_format)
	return new_time_date
}

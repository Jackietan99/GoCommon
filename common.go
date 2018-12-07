package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
* 传入数据,返回字符串类型
 */
func AssertionData(data interface{}) string {
	res := ""
	if value, ok := data.(string); ok {
		res = value
	} else if value, ok := data.([]byte); ok {
		res = string(value)
	} else if value, ok := data.(int); ok {
		res = strconv.Itoa(value)
	} else if value, ok := data.(int64); ok {
		res = strconv.FormatInt(int64(value), 10)
	} else if value, ok := data.(float64); ok {
		res = strconv.FormatFloat(value, 'f', -1, 64)
	} else if value, ok := data.(bool); ok {
		if value {
			res = "true"
		} else {
			res = "false"
		}
	}

	return res
}

/**
* 获取本机公网IP
 */
func GetNetworkIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return ""
	}
	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	res := string(contents)
	res = strings.Replace(res, "\n", "", -1)
	return res
}

/**
* 获取访问网站的客户端类型
* @userAgent string 浏览器中的user-Agent
* @return string 客户端类型
* pc:电脑端  android:安卓手机  iphone:iphone，ipad等移动端
 */
func GetWebVisitType(userAgent string) string {
	res := "pc"
	agent := strings.ToLower(userAgent)
	if strings.Contains(agent, "android") {
		res = "android"
	}

	if strings.Contains(agent, "iphone") {
		res = "iphone"
	}

	if strings.Contains(agent, "ipod") || strings.Contains(agent, "ipad") {
		res = "ipad"
	}

	return res
}

/**
* 通过淘宝IP地址库获取ip的详细信息
 */
func GetAddressByIP(ip string) map[string]string {
	//构造一个解析淘宝ip地址库接口返回的json格式
	type IpCode struct {
		Code interface{}
		Data map[string]interface{}
	}
	res := map[string]string{}
	url := "http://ip.taobao.com/service/getIpInfo.php"
	geturl := fmt.Sprintf("%s?ip=%s", url, ip)

	httpres, httpstatus := HttpRequest(geturl, "GET", "", "", "", nil)
	if httpstatus != 200 {
		return res
	}
	var jsonCode IpCode
	err := json.Unmarshal([]byte(httpres), &jsonCode)
	if err == nil {
		if AssertionData(jsonCode.Code) == "0" {
			res["country"] = AssertionData(jsonCode.Data["country"])
			res["country_id"] = AssertionData(jsonCode.Data["country_id"])
			res["area"] = AssertionData(jsonCode.Data["area"])
			res["area_id"] = AssertionData(jsonCode.Data["area_id"])
			res["region"] = AssertionData(jsonCode.Data["region"])
			res["region_id"] = AssertionData(jsonCode.Data["region_id"])
			res["city"] = AssertionData(jsonCode.Data["city"])
			res["city_id"] = AssertionData(jsonCode.Data["city_id"])
			res["county"] = AssertionData(jsonCode.Data["county"])
			res["county_id"] = AssertionData(jsonCode.Data["county_id"])
			res["isp"] = AssertionData(jsonCode.Data["isp"])
			res["isp_id"] = AssertionData(jsonCode.Data["isp_id"])
			res["ip"] = AssertionData(jsonCode.Data["ip"])
		}
	}
	return res
}

/**
* 正则匹配
* @regStr string 正则表达式
* @str string 字符串
* @return 字符串中匹配正则表达式的值，以切片的形式返回
 */
func RegMatch(regStr, str string) []string {
	reg := regexp.MustCompile(regStr)
	res := reg.FindAllString(str, -1)
	return res
}

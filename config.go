package common

import (
	"fmt"
	"os"

	"github.com/widuu/goini"
)

var conf *goini.Config
var inipath string //配置文件的路径

/**
* 手动指定config路径
* @path string 配置文件的路径
 */
func SetConf(path string) *goini.Config {
	inipath = "/etc/url_jobs.ini"
	//inipath = "config.ini"
	if len(path) > 0 {
		_, err := os.Stat(path)
		if err == nil {
			conf = goini.SetConfig(path)
		} else {
			fmt.Println("ini file is error")
			os.Exit(1)
		}
	} else {
		conf = goini.SetConfig(inipath)
	}
	return conf
}

/**
* 读取配置文件
* @dt string 读取配置文件中的[key] 其中的key
* @dl string 读取配置文件中的 field=value 中的field
* return string 返回值为配置文件中的value
 */
func GetConf(dt string, dl string) string {
	if conf == nil {
		return "config.in path error"
		os.Exit(1)
	}

	if len(conf.ReadList()) < 1 {
		return "config.ini content error"
		os.Exit(1)
	}
	res := conf.GetValue(dt, dl)
	return res
}

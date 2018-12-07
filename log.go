package common

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var Logpath = "/home/log/gamePool"

/**
* 设置日志目录
 */
func SetLogPath(path string) {
	Logpath = path
}

/**
* 指定文件名写入文件
* @logpath		日志目录
* @FiLeName		日志名称
 */
func LogsWithFileName(logpath, FiLeName string, msg string) {

	if len(logpath) < 1 {
		logpath = Logpath
	}

	if IsDirExists(logpath) != true {
		os.MkdirAll(logpath, 0777)
	}
	isEnd := strings.HasSuffix(logpath, "/")
	if isEnd != true {
		logpath = logpath + "/"
	}

	timeStr := time.Now().Format("2006-01-02")
	logFile := logpath + FiLeName + timeStr + ".log"

	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("create log err->", err)
		return
	}

	fout.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\r\n" + msg + "\r\n=====================\r\n")
	defer fout.Close()
}

/**
* 判断目录是否存在
 */
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

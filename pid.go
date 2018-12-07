/**
* unix/linux的pid文件写入，主要是为了确保进程唯一执行
* 原理:在/tmp/进程名称.pid  来确保一个进程只有一个pid
* 使用方法
	1, ChkRun(sExeName)		//判断进程是否正在运行中, 进程在运行中将强制退出，否则无返回执行下去
	2, CloseRun(sExeName)	//关闭进程
*/
package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

/**
* 关闭进程
* @param			sExeName		进程名称
* kill  pid
* 删除pid文件
 */
func CloseRun(sExeName string) int {
	pid := 0
	sPathFile := getPidPath(sExeName)
	_, err := os.Stat(sPathFile)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		f, _ := os.Open(sPathFile) //打开文件
		defer f.Close()
		buff := bufio.NewReader(f)       //读入缓存
		line, _ := buff.ReadString('\n') //以'\n'为结束符读入一行
		pid, _ := strconv.Atoi(line)

		//杀掉进程
		pro, err2 := os.FindProcess(pid)
		if err2 != nil {
			fmt.Println(line, pid, err2.Error())
		} else {
			err3 := pro.Kill()
			if err3 != nil {
				fmt.Println(err3.Error())
			} else {
				fmt.Println("app close done")
			}

		}

		//删除pid
		delPid(sPathFile)
		os.Exit(1)
	}
	return pid
}

/**
* 判断进程是否已经在运行，如果发现运行中，则不要启动本进程
* @param			sExeName			进程名称
 */
func ChkRun(sExeName string) bool {
	sPathFile := getPidPath(sExeName)
	_, err := os.Stat(sPathFile)
	if err != nil && os.IsNotExist(err) {
		setPid(sPathFile)
		return false
	} else {
		f, _ := os.Open(sPathFile) //打开文件
		defer f.Close()
		buff := bufio.NewReader(f)       //读入缓存
		line, _ := buff.ReadString('\n') //以'\n'为结束符读入一行
		pid, _ := strconv.Atoi(line)

		//_, err := os.FindProcess(pid)
		err = syscall.Kill(pid, 0)
		if err != nil {
			setPid(sPathFile)
		} else {
			//进程已经存在，强制关闭进程
			fmt.Println(sExeName, " is runing other process,pid->", pid)
			os.Exit(1)
		}
		//进程已经存在，强制关闭进程
		fmt.Println(sExeName, " is runing other process,pid->,", pid, "please run stop before")
	}
	return true
}

/**
* 获得pid位置
* @param		sExeName		写入进程的名称
 */
func getPidPath(sExeName string) string {
	path := "/tmp"
	sPathFile := fmt.Sprintf("%s/%s.pid", path, sExeName) //pid文件的完整路经

	//创建pid的目录
	if IsDirExists(path) != true {
		os.MkdirAll(path, 0777)
	}
	return sPathFile
}

/**
* 写入pid文件
* @param		sPathFile		写入pid
 */
func setPid(sPathFile string) {
	pid := strconv.Itoa(os.Getpid())
	//打开pid文件，并且写入pid
	fout, err := os.OpenFile(sPathFile, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(sPathFile, err)
		return
	}
	fout.WriteString(pid)
	defer fout.Close()
}

/**
* 删除pid文件
 */
func delPid(sPathFile string) {
	os.Remove(sPathFile)
}

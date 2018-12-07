package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/**
* 读取文件内容并返回字符串
* @param  path  文件路径
 */
func ReadFileString(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		//panic(err)
		return ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

/**
* 获取文件的修改时间
* 返回当前unix时间戳
 */
func GetFileModifyTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

/**
* 获取文件大小,单位时B
 */
func GetFileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

/**
* 创建文件
 */
func CreateFile(src string) (string, error) {
	//	src := dir + name + "/"
	if IsExist(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0777); err != nil {
		if os.IsPermission(err) {
			fmt.Println("你不够权限创建文件")
		}
		return "", err
	}

	return src, nil
}

/**
* 删除文件
 */
func DeleteFile(file string) error {
	return os.Remove(file)
}

/**
* 重命名文件
 */
func RenameFile(file string, to string) error {
	return os.Rename(file, to)
}

/**
* 新建文件并写入内容
* 如果文件已存在,则覆盖以前内容
 */
func WriteFile(filePath, fileName, content string) (int, error) {
	_, err := CreateFile(filePath)
	if err != nil {
		return 0, err
	}
	src := filePath + "/" + fileName
	fs, e := os.Create(src)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

/**
* 判断文件是否存在,读取文件内容
 */
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

/**
* 判断是否时文件
 */
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

/**
* 判断文件或者目录是否存在
 */
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

var localFileList []string

//d = append(d[:dLen], insertSlice[:insertSliceLen]...)
func walkFunc(path string, info os.FileInfo, err error) error {
	localFileList = append(localFileList, path)
	//fmt.Printf("%s\n", path)
	//fmt.Println("文件目录：", fileList)16
	return nil
}

/**
* 获取子目录和文件列表
 */
func GetFileList(localPath string) []string {
	var fileList []string
	filepath.Walk(localPath, walkFunc)
	fileList = localFileList
	localFileList = []string{}
	return fileList
}

/**
* 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤
* @dirPth string 目录路径
* @suffix string 匹配后缀过滤
 */
func GetAllFileName(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

/**
* 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤
* @dirPth string 目录路径
* @suffix string 匹配后缀过滤
 */
func GetAllWalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

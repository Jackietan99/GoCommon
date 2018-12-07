package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
* 获得程序的执行路径
 */
func GetExecPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dirctory, _ := filepath.Abs(file)

	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

package utils

import (
	"os/exec"
)

// 在指定的目录下运行命令
func RunCmdByDir(dir string, cmdName string, arg ...string) (string, error) {
	// exec.Command 返回一个新的Cmd
	cmd := exec.Command(cmdName, arg...)
	// 命令运行的目录
	cmd.Dir = dir
	// 标准输出和错误输出合并
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

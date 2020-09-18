package utils

import (
	"errors"
	"strings"
)

// 获取仓库名称
func GetRepoName(gitUrl string) (string, error) {

	// 判断url是否以.git 结尾
	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("git URL must end with .git！")
	}

	// 去除.git 结尾的url
	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")
	// 以 / 为分隔符 拿取截取的最后一个名称 为仓库的名称
	urlArr := strings.Split(noSuffixUrl, "/")

	return urlArr[len(urlArr)-1], nil
}

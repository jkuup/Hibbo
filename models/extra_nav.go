package models

import (
	"sort"
	"strings"
)

// 导航结构体
type Nav struct {
	// 标题
	Title string
	// 文章路径
	Path string
}
type Navs []Nav

func initExtraNav(dir string) (Navs, error) {

	var navigation Navs
	var extraNav Articles

	extraNav, err := RecursiveReadArticles(dir)
	if err != nil {
		return navigation, err
	}
	sort.Sort(extraNav)

	for _, article := range extraNav {
		title := strings.Title(strings.ToLower(article.Title))
		navigation = append(navigation, Nav{Title: title, Path: article.Path})
	}

	return navigation, nil
}

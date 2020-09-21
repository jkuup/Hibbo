package models

import (
	"math"
)

type PageResult struct {
	// 展示每页的文章个数，最多为 3
	List Articles `json:"list"`
	// 文章的总数
	Total int `json:"total"`
	// 展示第几页，比如10页，下面就显示 1 2 3 4 5
	Page int `json:"page"`
	// 每页显示的文章数 默认为 3
	PageSize int `json:"pageSize"`
	// 总共多少页 比如有10页 显示[10]
	TotalPage int
}

// 分页处理函数
func Pagination(articles *Articles, page int, pageSize int) PageResult {

	articleLen := len(*articles)
	// 每页显示3篇文章，页数就等于文章的个数/pageSize
	totalPage := int(math.Floor(float64(articleLen / pageSize)))

	if (articleLen % pageSize) != 0 {
		totalPage++
	}
	result := PageResult{
		Total:     articleLen,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}
	if page < 1 {
		result.Page = 1
	}

	// page 最大等于totalPage
	if page > result.TotalPage {
		result.Page = result.TotalPage
	}

	// 展示文章处理方法
	if articleLen <= result.PageSize {
		result.List = (*articles)[0:articleLen]
	} else {
		startNum := (result.Page - 1) * result.PageSize
		endNum := startNum + result.PageSize
		if endNum > articleLen {
			endNum = articleLen
		}
		// 每页展示的文章
		result.List = (*articles)[startNum:endNum]
	}

	return result
}

package main

import (
	"pibbo/models"
	"pibbo/route"
)

func init() {
	models.CompiledContent() //克隆或者更新文章、递归生成文章、导航、短链 Map
}

func main() {

	route.InitRoute()

}

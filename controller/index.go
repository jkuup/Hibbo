package controller

import (
	"net/http"
	"pibbo/config"
	"pibbo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//indexTemplate := models.Template.Index

	if err := c.Request.ParseForm(); err != nil {
		//indexTemplate.WriteError(c.Writer, err)
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}

	page, err := strconv.Atoi(c.Request.Form.Get("page"))
	if err != nil {
		page = 1
	}
	articles := models.ArticleList

	search := c.Request.Form.Get("search")
	category := c.Request.Form.Get("category")

	if search != "" || category != "" {
		articles = models.ArticleSearch(&articles, search, category)
	}

	result := models.Pagination(&articles, page, config.Cfg.PageSize)

	//indexTemplate.WriteData(c.Writer, models.BuildViewData("Blog", result))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":  "Blog",
		"Data":   result,
		"Config": config.Cfg,
		"Navs":   models.Navigation,
	})

}

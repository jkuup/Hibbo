package controller

import (
	"net/http"
	"pibbo/config"
	"pibbo/models"

	"github.com/gin-gonic/gin"
)

func Article(c *gin.Context) {
	//articleTemplate := models.Template.Article

	if err := c.Request.ParseForm(); err != nil {
		//articleTemplate.WriteError(c.Writer, err)
		c.HTML(http.StatusOK, "article.html", gin.H{})
	}
	key := c.Request.Form.Get("key")

	path := models.ArticleShortUrlMap[key]

	articleDetail, err := models.ReadArticleDetail(path)

	if err != nil {
		//articleTemplate.WriteError(c.Writer, err)

		c.HTML(http.StatusOK, "article.html", gin.H{})
	}

	//articleTemplate.WriteData(c.Writer, models.BuildViewData("Article", articleDetail))
	c.HTML(http.StatusOK, "article.html", gin.H{
		"Title":  "Article",
		"Data":   articleDetail,
		"Config": config.Cfg,
		"Navs":   models.Navigation,
	})
}

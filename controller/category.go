package controller

import (
	"net/http"
	"pibbo/config"
	"pibbo/models"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {
	//categoriesTemplate := models.Template.Categories

	result := models.GroupByCategory(&models.ArticleList, config.Cfg.CategoryDisplayQuantity)

	//categoriesTemplate.WriteData(c.Writer, models.BuildViewData("Blog", result))

	c.HTML(http.StatusOK, "categories.html", gin.H{
		"Title":  "Categories",
		"Data":   result,
		"Config": config.Cfg,
		"Navs":   models.Navigation,
	})
}

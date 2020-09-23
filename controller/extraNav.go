package controller

import (
	"net/http"
	"pibbo/config"
	"pibbo/models"

	"github.com/gin-gonic/gin"
)

func ExtraNav(c *gin.Context) {
	//extraNavTemplate := models.Template.ExtraNav

	if err := c.Request.ParseForm(); err != nil {
		//extraNavTemplate.WriteError(c.Writer, err)
		c.HTML(http.StatusOK, "extraNav.html", gin.H{})
	}

	name := c.Request.Form.Get("name")
	for _, nav := range models.Navigation {
		if nav.Title == name {
			articleDetail, err := models.ReadArticleDetail(nav.Path)
			if err != nil {
				//extraNavTemplate.WriteError(c.Writer, err)
				c.HTML(http.StatusOK, "extraNav.html", gin.H{})
			}
			//extraNavTemplate.WriteData(c.Writer, models.BuildViewData(nav.Title, articleDetail))
			c.HTML(http.StatusOK, "extraNav.html", gin.H{
				"Title":  nav.Title,
				"Data":   articleDetail,
				"Config": config.Cfg,
				"Navs":   models.Navigation,
			})
			return
		}
	}
}

package controller

import (
	"net/http"
	"pibbo/config"
	"pibbo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	var dashboardMsg []string
	//dashboardTemplate := models.Template.Dashboard

	if err := c.Request.ParseForm(); err != nil {
		//dashboardTemplate.WriteError(c.Writer, err)
		c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	}

	index, err := strconv.Atoi(c.Request.Form.Get("theme"))
	if err == nil && index < len(config.Cfg.ThemeOption) {
		config.Cfg.ThemeColor = config.Cfg.ThemeOption[index]
		dashboardMsg = append(dashboardMsg, "颜色切换成功!")
	}

	action := c.Request.Form.Get("action")
	if "updateArticle" == action {
		models.CompiledContent()
		dashboardMsg = append(dashboardMsg, "文章更新成功!")
	}

	/*
		dashboardTemplate.WriteData(c.Writer, models.BuildViewData("Dashboard", map[string]interface{}{
			"msg": dashboardMsg,
		}))
	*/

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Title": "Dashboard",
		"Data": map[string]interface{}{
			"msg": dashboardMsg,
		},
		"Config": config.Cfg,
		"Navs":   models.Navigation,
	})

}

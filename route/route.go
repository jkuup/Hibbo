package route

import (
	"html/template"
	"net/http"
	"pibbo/config"
	"pibbo/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"SpreadDigit": func(n int) []int {
			var r []int
			for i := 1; i <= n; i++ {
				r = append(r, i)
			}
			return r
		},
	})
	r.LoadHTMLGlob("views/*")
	r.StaticFS("/public/", http.Dir(config.Cfg.CurrentDir+"/public"))
	r.StaticFS("/assets/", http.Dir(config.Cfg.DocumentAssetsDir))
	r.StaticFS("/images/", http.Dir(config.Cfg.CurrentDir+"/images"))
	r.GET("/", controller.Index)
	r.GET("/blog", controller.Index)
	r.GET("/categories", controller.Category)
	r.GET("/article", controller.Article)
	r.GET("/extra-nav", controller.ExtraNav)

	r.GET(config.Cfg.GitHookUrl, controller.GithubHook)
	r.GET(config.Cfg.DashboardEntrance, controller.Dashboard)
	r.Run(":8080")
}

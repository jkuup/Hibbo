package controller

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"pibbo/config"
	"pibbo/models"

	"github.com/gin-gonic/gin"
)

func GithubHook(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		SedResponse(c, err.Error())
		return
	}

	if "" == config.Cfg.WebHookSecret || "push" != c.Request.Header.Get("x-github-event") {
		SedResponse(c, "No Configuration WebHookSecret Or Not Pushing Events")
		log.Println("No Configuration WebHookSecret Or Not Pushing Events")
		return
	}

	sign := c.Request.Header.Get("X-Hub-Signature")

	bodyContent, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		SedResponse(c, err.Error())
		log.Println("WebHook err:" + err.Error())
		return
	}

	if err = c.Request.Body.Close(); err != nil {
		SedResponse(c, err.Error())
		log.Println("WebHook err:" + err.Error())
		return
	}
	mac := hmac.New(sha1.New, []byte(config.Cfg.WebHookSecret))
	mac.Write(bodyContent)
	expectedHash := "sha1=" + hex.EncodeToString(mac.Sum(nil))

	if sign != expectedHash {
		SedResponse(c, "WebHook err:Signature does not match")
		log.Println("WebHook err:Signature does not match")
		return
	}

	SedResponse(c, "ok")

	models.CompiledContent()

}

func SedResponse(c *gin.Context, msg string) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := c.Writer.Write([]byte(`{"msg": "` + msg + `"}`))
	if err != nil {
		log.Println(err)
	}
}

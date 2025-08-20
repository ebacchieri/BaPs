package sdk

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/config"
)

type SDK struct {
	router *gin.Engine
}

func New(router *gin.Engine) *SDK {
	s := &SDK{
		router: router,
	}

	s.initRouter()
	return s
}

func (s *SDK) initRouter() {
	s.router.LoadHTMLGlob(fmt.Sprintf("%s/templates/*", config.GetConfig().DataPath))
	s.router.Any("/", handleIndex)
	s.router.Any("/index", handleIndex)

	s.router.GET("/r:url", s.connectionGroups)
	s.router.GET("/prod/index.json", index)

	account := s.router.Group("/account")
	{
		account.POST("/yostar_auth_request", s.YostarAuthRequest)
		account.POST("/yostar_auth_submit", s.YostarAuthSubmit)
	}
	user := s.router.Group("/user")
	{
		user.POST("/yostar_createlogin", s.YostarCreatelogin)
		user.POST("/login", s.YostarLogin)
		//user.POST("/transcode_verify", s.TranscodeVerify)
		user.Any("/agreement", agreement)
	}
	app := s.router.Group("/app")
	{
		app.Any("/getSettings", getSettings)
		app.Any("/getCode", getCode)
	}
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Ba Ps!",
		"github": "https://github.com/gucooing/BaPs",
	})
}

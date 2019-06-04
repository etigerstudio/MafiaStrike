package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(g *gin.Engine) {
	g.Static("statics", "statics")

	g.GET("",  func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
		})
	})
}
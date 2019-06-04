package server

import (
	"github.com/gin-gonic/gin"
	"mafia-strike/controllers"
	"net/http"
)

func initRouter(g *gin.Engine) {
	g.Static("statics", "statics")

	g.GET("",  func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	lobby := g.Group("lobbies")
	{
		lobby.POST("", controllers.PostLobbyEntry)
		lobby.GET(":lobby", controllers.GetLobbyEntry)
	}
}
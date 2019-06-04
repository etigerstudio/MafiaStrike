package server

import (
	"github.com/gin-gonic/gin"
	"mafia-strike/util"
)

func Run() {
	util.InitRandSeed()
	g := gin.Default()
	g.LoadHTMLGlob("templates/*")
	initRouter(g)
	_ = g.Run("0.0.0.0:8888")
}
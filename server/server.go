package server

import (
	"github.com/gin-gonic/gin"
	"mafia-strike/util"
)

func Run() {
	util.InitRandSeed()
	g := gin.Default()
	initRouter(g)
	_ = g.Run("0.0.0.0:6000")
}
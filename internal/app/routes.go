package app

import (
	"rank/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine) {
	router.GET("/healthz", handler.HandleHealthz)
	router.POST("/match/score", handler.AddScore)
	router.GET("/match/rank", handler.GetRank)
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(gameHandler *GameHandler) *gin.Engine {
	router := gin.New()

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/api/v1/leaderboard", gameHandler.GetLeaderboard)
	router.POST("/api/v1/game", gameHandler.AddGame)

	return router
}

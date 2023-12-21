package main

import (
	"go-gin-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/player", handlers.GetPlayerInfoHandler)
	r.POST("/player", handlers.AddPlayerHandler)
	r.DELETE("/player/:id", handlers.DeletePlayerHandler)
	r.PATCH("/player/:id", handlers.UpdatePlayerInfoHandler)
	r.GET("/leaderboard", handlers.GetLeaderboardHandler)

	r.Run(":8080")
}
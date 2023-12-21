package handlers

import (
	"go-gin-rest-api/entities"
	"go-gin-rest-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeaderboardHandler(c *gin.Context) {
	response, err := controllers.GetLeaderboard(entities.TableName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": response,
	})
}
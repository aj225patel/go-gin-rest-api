package handlers

import (
	"go-gin-rest-api/entities"
	"go-gin-rest-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPlayerHandler(c *gin.Context) {
	var player entities.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controllers.InsertNewPlayer(entities.TableName, player.PlayerId, player.PlayerName, player.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "player has been successfully added",
	})
}
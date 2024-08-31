package handlers

import (
    "github.com/parkrealgood/gotification/models"
    "github.com/parkrealgood/gotification/services"
		"github.com/parkrealgood/gotification/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)


func CreateTopic(c *gin.Context) {
    var newTopic models.Topic
    if err := c.ShouldBindJSON(&newTopic); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", "INVALID_REQUEST", err.Error())
			return
    }
    topic, err := services.CreateTopic(&newTopic)
    if err != nil {
				utils.RespondWithError(c, http.StatusInternalServerError, "Topic Create Error", "CREATE_ERROR", err.Error())
        return
    }
    c.JSON(http.StatusOK, topic)
}

func GetTopics(c *gin.Context) {
		topics := services.GetTopics()
		c.JSON(http.StatusOK, topics)
}

func GetTopic(c *gin.Context) {
		id := c.Param("id")
		topic, err := services.GetTopic(id)
		if err != nil {
				utils.RespondWithError(c, http.StatusNotFound, "Topic Not Found", "NOT_FOUND", err.Error())
				return
		}
		c.JSON(http.StatusOK, topic)
}

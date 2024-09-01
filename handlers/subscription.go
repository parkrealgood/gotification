package handlers

import (
	"net/http"

	"sync"

	"github.com/gin-gonic/gin"
	"github.com/parkrealgood/gotification/services"
	"github.com/parkrealgood/gotification/utils"
)

func SubscribeTopic(c *gin.Context) {
	topicID := c.Param("id")
	var request struct {
		UserID string `json:"UserID" binding:"required"`
	}

	if errBindJSON := c.ShouldBindJSON(&request); errBindJSON != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "INVALID_REQUEST", errBindJSON.Error())
		return
	}

	_, errGetTopic := services.GetTopic(topicID)
	if errGetTopic != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Topic not found", "NOT_FOUND", errGetTopic.Error())
		return
	}

	subscription, errSubscribeTopic := services.SubscribeTopic(request.UserID, topicID)
	if errSubscribeTopic != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Subscription Error", "SUBSCRIPTION_ERROR", errSubscribeTopic.Error())
		return
	}
	c.JSON(http.StatusOK, subscription)
}

func PublishTopic(c *gin.Context) {
	topicID := c.Param("id")

	var request struct {
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "INVALID_REQUEST", err.Error())
		return
	}

	subscribers := services.GetTopicSubscribers(topicID)
	if len(subscribers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No subscribers for this topic"})
		return
	}

	var wg sync.WaitGroup
	for _, userID := range subscribers {
		wg.Add(1)
		go func(uid string) {
			defer wg.Done()
			services.SendMessageToUser(uid, request.Message, topicID)
		}(userID)
	}

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "Message sent to all subscribers"})
}

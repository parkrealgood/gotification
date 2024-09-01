package handlers

import (
	"net/http"

	"sync"

	"github.com/gin-gonic/gin"
	"github.com/parkrealgood/gotification/models"
	"github.com/parkrealgood/gotification/services"
	"github.com/parkrealgood/gotification/utils"
)

var (
	topics        = make(map[string]*models.Topic)        // 토픽 데이터 저장
	subscriptions = make(map[string]*models.Subscription) // 구독 관계 저장
)

func SubscribeTopic(c *gin.Context) {
	topicID := c.Param("id") // URL에서 토픽 ID 추출
	var request struct {
		UserID string `json:"UserID" binding:"required"`
	}

	if errBindJSON := c.ShouldBindJSON(&request); errBindJSON != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "INVALID_REQUEST", errBindJSON.Error())
		return
	}

	_, errGetTopic := services.GetTopic(topicID)
	// _, errGetUser := services.GetUser(request.UserID)

	// if errGetUser != nil {
	// utils.RespondWithError(c, http.StatusNotFound, "User not found", "NOT_FOUND", errGetUser.Error())
	// return
	// }
	if errGetTopic != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Topic not found", "NOT_FOUND", errGetTopic.Error())
		return
	}

	// 구독 관계 생성
	subscription, errSubscribeTopic := services.SubscribeTopic(request.UserID, topicID)
	if errSubscribeTopic != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Subscription Error", "SUBSCRIPTION_ERROR", errSubscribeTopic.Error())
		return
	}
	c.JSON(http.StatusOK, subscription)
}

// 토픽 발행 핸들러
func PublishTopic(c *gin.Context) {
	topicID := c.Param("id")

	// 발행할 메시지 내용을 요청으로부터 가져옴
	var request struct {
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// 토픽 존재 여부 확인
	topic, topicExists := topics[topicID]
	if !topicExists {
		utils.RespondWithError(c, http.StatusNotFound, "Topic Not Found", "NOT_FOUND", "")
		return
	}

	// 해당 토픽을 구독한 유저 목록 조회
	var subscribers []string
	for _, sub := range subscriptions {
		if sub.TopicID == topicID {
			subscribers = append(subscribers, sub.UserID)
		}
	}

	if len(subscribers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No subscribers for this topic"})
		return
	}

	// 동시적으로 메시지를 보내기 위해 WaitGroup 사용
	var wg sync.WaitGroup
	for _, userID := range subscribers {
		wg.Add(1)
		go func(uid string) {
			defer wg.Done()
			services.SendMessageToUser(uid, request.Message, topic.Name)
		}(userID)
	}

	// 모든 메시지 전송이 완료될 때까지 대기
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "Message sent to all subscribers"})
}

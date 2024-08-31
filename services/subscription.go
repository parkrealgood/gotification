package services

import (
	"fmt"
	"time"

	"github.com/parkrealgood/gotification/models"
)

var subscriptions = make(map[string]*models.Subscription)

func SubscribeTopic(userId string, topicId string) (*models.Subscription, error) {
	// 구독 관계 생성
	subscriptionKey := userId + ":" + topicId
	if subscription, exists := subscriptions[subscriptionKey]; exists {
		return subscription, nil
	}

	subscriptions[subscriptionKey] = &models.Subscription{
		UserID:       userId,
		TopicID:      topicId,
		SubscribedAt: time.Now(),
	}
	return subscriptions[subscriptionKey], nil
}

// 유저에게 메시지 보내기 (예시로 출력)
func SendMessageToUser(userID string, message string, topicName string) {
	user, exists := users[userID]
	if !exists {
		fmt.Printf("User %s not found\n", userID)
		return
	}

	// 실제로 메시지를 보내는 로직을 여기에 구현
	fmt.Printf("Sending message to %s: [%s] %s\n", user.Name, topicName, message)
}

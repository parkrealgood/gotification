package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/parkrealgood/gotification/models"
)

var subscriptions = make(map[string]*models.Subscription)
var lastSubscriptionID int

func SubscribeTopic(userId string, topicId string) (*models.Subscription, error) {
	// 구독 관계 생성
	subscriptionKey := userId + ":" + topicId
	if subscription, exists := subscriptions[subscriptionKey]; exists {
		return subscription, nil
	}
	id := GenerateSubscriptionID()
	newSubscription := &models.Subscription{
		ID:           id,
		Key:          subscriptionKey,
		UserID:       userId,
		TopicID:      topicId,
		SubscribedAt: time.Now(),
	}
	subscriptions[subscriptionKey] = newSubscription
	return newSubscription, nil
}

func SendMessageToUser(userID string, message string, topicId string) {
	// 메시지 전송 시간을 1초로 가정
	time.Sleep(1 * time.Second)
	// 실제로 메시지를 보내는 로직을 여기에 구현
	fmt.Printf("Sending message to %s: [%s] %s\n", userID, topicId, message)
}

func GenerateSubscriptionID() string {
	idMutex.Lock()
	defer idMutex.Unlock()

	lastSubscriptionID++
	return strconv.Itoa(lastSubscriptionID)
}

func GetTopicSubscribers(topicID string) []string {
	var subscribers []string
	for _, sub := range subscriptions {
		if sub.TopicID == topicID {
			subscribers = append(subscribers, sub.UserID)
		}
	}
	return subscribers
}

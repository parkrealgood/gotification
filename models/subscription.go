package models

import "time"

type Subscription struct {
	ID           string    `json:"id"`
	TopicID      string    `json:"topic_id"`
	UserID       string    `json:"user_id"`
	SubscribedAt time.Time `json:"subscribed_at"`
}

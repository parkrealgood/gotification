package services

import (
    "github.com/parkrealgood/gotification/models"
    "errors"
)

var topics = make(map[string]*models.Topic)

func CreateTopic(topic *models.Topic) (*models.Topic, error) {
    if _, exists := topics[topic.ID]; exists {
        return nil, errors.New("topic already exists")
    }
    topics[topic.ID] = topic
    return topic, nil
}

func GetTopics() []*models.Topic {
		var topicList []*models.Topic
		for _, topic := range topics {
				topicList = append(topicList, topic)
		}
		return topicList 
}

func GetTopic(id string) (*models.Topic, error) {
		topic, exists := topics[id]
		if !exists {
				return nil, errors.New("topic not found")
		}
		return topic, nil
}

package services

import (
    "github.com/parkrealgood/gotification/models"
    "errors"
)

var users = make(map[string]*models.User)

func GetUser(id string) (*models.User, error) {
		user, exists := users[id]
		if !exists {
				return nil, errors.New("user not found")
		}
		return user, nil
}
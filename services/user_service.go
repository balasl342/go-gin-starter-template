package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/balasl342/go-gin-starter-template/models"
)

var users = map[string]models.User{}

func GetUserById(id string) (*models.User, error) {
	if user, exists := users[id]; exists {
		return &user, nil
	}
	return nil, errors.New("user not found")
}

func CreateUser(user *models.User) error {
	user.ID = generateID()
	user.CreatedAt = time.Now()
	users[user.ID] = *user
	return nil
}

func generateID() string {
	return uuid.New().String()
}

package dao

import (
	"github.com/A-Siam/go_auth/src/data"
	"github.com/A-Siam/go_auth/src/data/entities"
	"github.com/google/uuid"
)

func GetUser(username string) (entities.User, error) {
	user := entities.User{}
	err := data.GetDB().Where("username = ?", username).First(&user).Error
	return user, err
}

func SaveUser(username, password, groups string) (entities.User, error) {
	user := entities.User{ID: uuid.NewString(), Username: username, Password: password, Groups: groups}
	result := data.GetDB().Create(&user)
	return user, result.Error
}

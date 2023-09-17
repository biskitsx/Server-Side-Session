package service

import (
	"errors"
	"fmt"

	"github.com/biskitsx/Server-Side-Session/database"
	"github.com/biskitsx/Server-Side-Session/model"
)

type AuthService interface {
	CheckUsername(username string) (*model.User, error)
}
type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (authService) CheckUsername(username string) (*model.User, error) {
	user := &model.User{}
	database.Db.Where("username = ?", username).First(user)
	if user.Username != "" {
		fmt.Println(user)
		return user, nil
	}
	return user, errors.New("user not found")
}

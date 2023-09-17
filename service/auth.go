package service

import (
	"errors"
	"fmt"

	"github.com/biskitsx/Server-Side-Session/container"
	"github.com/biskitsx/Server-Side-Session/model"
)

type AuthService interface {
	CheckUsername(username string) (*model.User, error)
}
type authService struct {
	container container.Container
}

func NewAuthService(c container.Container) AuthService {
	return &authService{
		container: c,
	}
}

func (service *authService) CheckUsername(username string) (*model.User, error) {
	user := &model.User{}
	db := service.container.GetDatabase()
	db.Where("username = ?", username).First(user)
	if user.Username != "" {
		fmt.Println(user)
		return user, nil
	}
	return user, errors.New("user not found")
}

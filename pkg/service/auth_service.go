package service

import (
	"crypto/sha256"
	"fmt"
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
)

type AuthService struct {
	repo repostiroty.Authorization
}

func NewAuthService(repo repostiroty.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a AuthService) CreateUser(user todo_demo.User) (int, error) {
	user.Password = encryptPassword(user.Username, user.Password)
	return a.repo.CreateUser(user)
}

func encryptPassword(username, pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(username)))
}

package service

import (
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
)

type Authorization interface {
	CreateUser(user todo_demo.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repostiroty.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}

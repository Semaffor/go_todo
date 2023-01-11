package service

import "github.com/semaffor/go-todo-app/pkg/repostiroty"

type Authorization interface {
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
	return &Service{}
}

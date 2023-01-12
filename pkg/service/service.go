package service

import (
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
)

type Authorization interface {
	CreateUser(user todo_demo.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseJwt(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_demo.TodoList) (int, error)
	GetAll(userId int) ([]todo_demo.TodoList, error)
	GetById(userId, listId int) (todo_demo.TodoList, error)
	DeleteById(userId, listId int) (uint8, error)
	Update(userId, listId int, input todo_demo.UpdateListInput) error
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
		TodoList:      NewTodoListService(repo.TodoList),
	}
}

package repostiroty

import (
	"github.com/jmoiron/sqlx"
	todo_demo "github.com/semaffor/go-todo-app"
)

type Authorization interface {
	CreateUser(user todo_demo.User) (int, error)
	GetUser(login, passwordHash string) (todo_demo.User, error)
}

type TodoList interface {
	CreateList(userId int, list todo_demo.TodoList) (int, error)
	GetAll(userId int) ([]todo_demo.TodoList, error)
	GetById(userId, listId int) (todo_demo.TodoList, error)
	DeleteById(userId, listId int) (uint8, error)
	Update(userId, listId int, input todo_demo.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		TodoList:      NewTodoListRepo(db),
	}
}

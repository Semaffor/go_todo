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
	}
}

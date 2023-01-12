package service

import (
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
)

type TodoListService struct {
	repo repostiroty.TodoList
}

func NewTodoListService(repo repostiroty.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list todo_demo.TodoList) (int, error) {
	return t.repo.CreateList(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]todo_demo.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, listId int) (todo_demo.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) DeleteById(userId, listId int) (uint8, error) {
	return t.repo.DeleteById(userId, listId)
}

func (t *TodoListService) Update(userId, listId int, input todo_demo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return t.repo.Update(userId, listId, input)
}

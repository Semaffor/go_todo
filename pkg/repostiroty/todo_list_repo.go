package repostiroty

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_demo "github.com/semaffor/go-todo-app"
	"log"
	"strings"
)

type TodoListRepo struct {
	db *sqlx.DB
}

func NewTodoListRepo(db *sqlx.DB) *TodoListRepo {
	return &TodoListRepo{db: db}
}

func (r *TodoListRepo) CreateList(userId int, list todo_demo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	var listId int
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListRepo) GetAll(userId int) ([]todo_demo.TodoList, error) {
	var lists []todo_demo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListRepo) GetById(userId, listId int) (todo_demo.TodoList, error) {
	var list todo_demo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListRepo) DeleteById(userId, listId int) (uint8, error) {
	// using - additional tables for validation in block 'where'
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListTable, usersListsTable)
	executedAction, err := r.db.Exec(query, userId, listId)
	if err != nil {
		return 0, err
	}
	affected, err := executedAction.RowsAffected()

	if affected != 1 {
		return 204, nil
	}
	return 0, nil
}

func (r *TodoListRepo) Update(userId, listId int, input todo_demo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	log.Printf("updateQuery: %s\n args: %s", query, args)

	_, err := r.db.Exec(query, args...)
	return err
}

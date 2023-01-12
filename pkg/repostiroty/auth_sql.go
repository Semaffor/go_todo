package repostiroty

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_demo "github.com/semaffor/go-todo-app"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (u *AuthRepo) CreateUser(user todo_demo.User) (int, error) {
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", usersTable)
	row := u.db.QueryRowx(query, user.Name, user.Username, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

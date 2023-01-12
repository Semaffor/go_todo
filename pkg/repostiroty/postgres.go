package repostiroty

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_lists"
	usersListsTable = "users_lists"
	todoItemTable   = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

/*
	.Open() - need additional check err on .Ping() manually,
	.Connect() check ping automatically // panic?
*/

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SslMode))

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

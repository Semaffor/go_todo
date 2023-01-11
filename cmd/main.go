package main

import (
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/handler"
	"log"
)

func main() {
	ginHandlers := handler.Handler{}

	srv := new(todo_demo.Server)
	if err := srv.Run("8080", ginHandlers.InitRoutes()); err != nil {
		log.Fatalf("Error when running: %s", err.Error())
	}
}

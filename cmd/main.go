package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/handler"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
	"github.com/semaffor/go-todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("App starting...")
	if err := initConfig(); err != nil {
		log.Fatalf("Error occurred: init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error with loading env: %s", err.Error())
	}

	db, err := repostiroty.NewPostgresDb(repostiroty.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
		SslMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error occurred when conecting to db: %s", err.Error())
	}

	repo := repostiroty.NewRepository(db)
	services := service.NewService(repo)
	ginHandlers := handler.NewHandler(services)

	srv := new(todo_demo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), ginHandlers.InitRoutes()); err != nil {
			log.Fatalf("Error when running: %s", err.Error())
		}
	}()

	controlChan := make(chan os.Signal, 1)
	signal.Notify(controlChan, syscall.SIGTERM, syscall.SIGINT)
	<-controlChan

	log.Println("App closing...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s\n", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

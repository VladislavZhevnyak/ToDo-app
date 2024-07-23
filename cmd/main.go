package main

import (
	"log"

	todo "github.com/VladislavZhevnyak/ToDo-app"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/handler"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/repository"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutees()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

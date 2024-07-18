package main

import (
	"log"

	todo "github.com/VladislavZhevnyak/ToDo-app"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutees()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

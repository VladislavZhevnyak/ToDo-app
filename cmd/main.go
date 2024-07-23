package main

import (
	"log"

	todo "github.com/VladislavZhevnyak/ToDo-app"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/handler"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/repository"
	"github.com/VladislavZhevnyak/ToDo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initilization configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutees()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

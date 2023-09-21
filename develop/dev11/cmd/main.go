package main

import (
	"log"
	"my_httpServer/cmd/service"
	"my_httpServer/cmd/service/repositories"
	"my_httpServer/config"
)

func main() {
	serviceConfig := config.GetConfig("config/config.yaml")

	eventRepository := repositories.NewEventRepository()
	userRepository := repositories.NewUserRepository()

	calendarService := service.New(eventRepository, userRepository, serviceConfig)

	if err := calendarService.Run(); err != nil {
		log.Fatalln(err)
	}
}

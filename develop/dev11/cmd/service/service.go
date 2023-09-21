package service

import (
	"log"
	"my_httpServer/cmd/service/repositories"
	"my_httpServer/config"
	"net/http"
)

type Service struct {
	eventRepository *repositories.EventRepository
	userRepository  *repositories.UserRepository
	config          config.Config
}

func New(
	eventRepository *repositories.EventRepository,
	userRepository *repositories.UserRepository,
	config config.Config) *Service {
	return &Service{
		eventRepository: eventRepository,
		userRepository:  userRepository,
		config:          config,
	}
}

func (s *Service) Run() error {
	server := &http.Server{
		Addr:    s.config.Service.Host + ":" + s.config.Service.Port,
		Handler: s.NewRouter(),
	}

	log.Println("Run calendar service:")
	log.Println("host:", s.config.Service.Host, "port:", s.config.Service.Port)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Service) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	loggerMiddleware := logger

	router.HandleFunc("/create_event", loggerMiddleware(s.CreateEvent))
	router.HandleFunc("/update_event", loggerMiddleware(s.UpdateEvent))
	router.HandleFunc("/delete_event", loggerMiddleware(s.DeleteEvent))

	router.HandleFunc("/events_for_day", loggerMiddleware(s.EventsForDay))
	router.HandleFunc("/events_for_week", loggerMiddleware(s.EventsForWeek))
	router.HandleFunc("/events_for_month", loggerMiddleware(s.EventsForMonth))

	return router
}

package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"my_httpServer/cmd/service/api"
	"my_httpServer/cmd/service/entities"
	"net/http"
	"time"
)

func (s *Service) CreateEvent(writer http.ResponseWriter, request *http.Request) {
	event, err := s.getEventHttpPostCreate(request)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.eventRepository.CreateEvent(event.UserID, *event)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := struct {
		Result uuid.UUID `json:"result"`
	}{event.ID}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = writer.Write(input)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadGateway)
	}
}

func (s *Service) UpdateEvent(writer http.ResponseWriter, request *http.Request) {
	event, err := s.getEventHttpPostUpdate(request)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.eventRepository.UpdateEvent(event.UserID, *event)

	result := struct {
		Result uuid.UUID `json:"result"`
	}{event.ID}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = writer.Write(input)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadGateway)
	}
}

func (s *Service) DeleteEvent(writer http.ResponseWriter, request *http.Request) {
	event, err := s.getEventHttpPostUpdate(request)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.eventRepository.DeleteEvent(event.UserID, event.ID)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := struct {
		Result uuid.UUID `json:"result"`
	}{event.ID}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = writer.Write(input)
	if err != nil {
		writeError(writer, err.Error(), http.StatusBadGateway)
	}
}

func (s *Service) getEventHttpPostCreate(r *http.Request) (*entities.Event, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("invalid method: " + r.Method)
	}

	dtoCreateEvent := api.DtoCreateEvent{}
	err := json.NewDecoder(r.Body).Decode(&dtoCreateEvent)
	if err != nil {
		return nil, err
	}

	date, err := parseDate(dtoCreateEvent.Event.Date)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(dtoCreateEvent.UserID)
	if err != nil {
		return nil, err
	}

	event := entities.Event{
		ID:     uuid.New(),
		Date:   date,
		UserID: userID,
		Info:   dtoCreateEvent.Event.Info,
	}

	return &event, nil
}

func (s *Service) getEventHttpPostUpdate(r *http.Request) (*entities.Event, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("invalid method: " + r.Method)
	}

	dtoCreateEvent := api.DtoCreateEvent{}
	err := json.NewDecoder(r.Body).Decode(&dtoCreateEvent)
	if err != nil {
		return nil, err
	}

	date, err := parseDate(dtoCreateEvent.Event.Date)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(dtoCreateEvent.UserID)
	if err != nil {
		return nil, err
	}

	eventId, err := uuid.Parse(dtoCreateEvent.Event.EventID)
	if err != nil {
		return nil, err
	}

	event := entities.Event{
		ID:     eventId,
		Date:   date,
		UserID: userID,
		Info:   dtoCreateEvent.Event.Info,
	}

	return &event, nil
}

func (s *Service) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, userID, err := UrlParams(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events := s.eventRepository.GetEventsForMonth(userID, *date)

	userEvents := entities.UserEvents{
		UserID: userID,
		Events: events,
	}

	result := struct {
		Result entities.UserEvents `json:"result"`
	}{userEvents}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = w.Write(input)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadGateway)
	}
}

func (s *Service) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, userID, err := UrlParams(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events := s.eventRepository.GetEventsForWeek(userID, *date)

	userEvents := entities.UserEvents{
		UserID: userID,
		Events: events,
	}

	result := struct {
		Result entities.UserEvents `json:"result"`
	}{userEvents}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = w.Write(input)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadGateway)
	}
}

func (s *Service) EventsForDay(w http.ResponseWriter, r *http.Request) {
	date, userID, err := UrlParams(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events := s.eventRepository.GetEventsOnDay(userID, *date)

	userEvents := entities.UserEvents{
		UserID: userID,
		Events: events,
	}

	result := struct {
		Result entities.UserEvents `json:"result"`
	}{userEvents}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err = w.Write(input)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadGateway)
	}
}

func UrlParams(r *http.Request) (*time.Time, uuid.UUID, error) {
	date, err := parseDate(r.URL.Query().Get("date"))
	if err != nil {
		return nil, uuid.Nil, err
	}
	parse, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		return nil, uuid.Nil, err
	}

	return &date, parse, nil
}

func parseDate(t string) (time.Time, error) {
	const dateLayout = "2006-01-02"
	timeReq, err := time.Parse(dateLayout, t)
	return timeReq, err
}

func writeError(w http.ResponseWriter, msg string, statusCode int) {
	err := struct {
		Err string `json:"error"`
	}{Err: msg}

	input, _ := json.MarshalIndent(&err, "", "  ")
	http.Error(w, string(input), statusCode)
}

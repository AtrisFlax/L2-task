package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"my_httpServer/cmd/service/entities"
	"sync"
	"time"
)

type EventRepository struct {
	events map[uuid.UUID]entities.Event
	sync.RWMutex
}

func NewEventRepository() *EventRepository {
	events := make(map[uuid.UUID]entities.Event)
	return &EventRepository{events: events}
}

func (er *EventRepository) CreateEvent(userUUID uuid.UUID, event entities.Event) (*entities.Event, error) {
	er.Lock()
	defer er.Unlock()

	if event.UserID != userUUID {
		return nil, fmt.Errorf("no such user_id: %s ", userUUID)
	}
	er.events[event.ID] = event
	return &event, nil
}

func (er *EventRepository) DeleteEvent(userUUID uuid.UUID, eventID uuid.UUID) error {
	er.Lock()
	defer er.Unlock()

	event, ok := er.events[eventID]
	if !ok {
		return fmt.Errorf("can't find event with %s event_id", eventID.String())
	}
	if event.UserID != userUUID {
		return fmt.Errorf("no such event_id: %s for user_id: %s", eventID.String(), userUUID)
	}

	delete(er.events, eventID)
	return nil
}

func (er *EventRepository) UpdateEvent(userUUID uuid.UUID, updatedEvent entities.Event) error {
	er.Lock()
	defer er.Unlock()

	_, ok := er.events[updatedEvent.ID]
	if !ok {
		return fmt.Errorf("can't find updatedEvent with %s event_id", updatedEvent.ID.String())
	}
	if updatedEvent.UserID != userUUID {
		return fmt.Errorf("no such event_id: %s for user_id: %s", updatedEvent.ID.String(), userUUID)
	}

	er.events[updatedEvent.ID] = updatedEvent
	fmt.Printf("INSIDE UpdateEvent   %+v\n", er.events)
	return nil
}

func (er *EventRepository) GetEvents(userUUID uuid.UUID) []entities.Event {
	er.RLock()
	defer er.RUnlock()

	var userEvents []entities.Event
	for _, event := range er.events {
		if event.UserID == userUUID {
			userEvents = append(userEvents, event)
		}
	}
	return userEvents
}

func (er *EventRepository) GetEventsOnDay(userUUID uuid.UUID, time time.Time) []entities.Event {
	er.RLock()
	defer er.RUnlock()

	var userEvents []entities.Event
	for _, event := range er.events {
		if event.UserID == userUUID && time == event.Date {
			userEvents = append(userEvents, event)
		}
	}
	return userEvents
}

func (er *EventRepository) GetEventsForWeek(userUUID uuid.UUID, targetTime time.Time) []entities.Event {
	er.RLock()
	defer er.RUnlock()
	targetYear, targetWeek := targetTime.ISOWeek()
	var userEvents []entities.Event
	for _, event := range er.events {
		year, week := event.Date.ISOWeek()
		if event.UserID == userUUID && targetYear == year && targetWeek == week {
			userEvents = append(userEvents, event)
		}
	}
	return userEvents
}

func (er *EventRepository) GetEventsForMonth(userUUID uuid.UUID, targetTime time.Time) []entities.Event {
	er.RLock()
	defer er.RUnlock()

	targetYear, targetMonth := targetTime.Year(), targetTime.Month()
	var userEvents []entities.Event
	for _, event := range er.events {
		year, month := event.Date.Year(), event.Date.Month()
		if event.UserID == userUUID && targetYear == year && targetMonth == month {
			userEvents = append(userEvents, event)
		}
	}
	return userEvents
}

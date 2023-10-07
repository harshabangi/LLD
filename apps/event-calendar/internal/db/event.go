package db

import (
	"github.com/harshabangi/LLD/apps/event-calendar/internal/models"
	"time"
)

type eventStorage struct {
	events map[time.Time][]models.Event
}

type EventDB interface {
	CreateEvent(event *models.Event)
	GetEvents(date time.Time) []models.Event
}

func NewEventDB() EventDB {
	return &eventStorage{}
}

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (e *eventStorage) CreateEvent(event *models.Event) {
	e.events[event.Date] = append(e.events[event.Date], *event)
}

func (e *eventStorage) GetEvents(date time.Time) []models.Event {
	return e.events[date]
}

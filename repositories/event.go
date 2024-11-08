package repositories

import (
	"context"
	"time"

	"github.com/ticket-booking-app/models"
)

type EventRepository struct {
	db any // hypothetical database connection
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		Id: "001",
		Name: "Comic Con",
		Location: "Gurgram",
		Date: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
    // database logic here to get a specific event by its ID
    return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
    // database logic here to create a new event
    return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

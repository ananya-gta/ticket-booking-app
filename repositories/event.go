package repositories

import (
	"context"
	"github.com/ticket-booking-app/models"
)

type EventRepository struct {
	db any // hypothetical database connection
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	return nil, nil
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

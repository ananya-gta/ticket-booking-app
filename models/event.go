package models

import (
	"time"
	"context"
)

type Event struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//  define methods for interacting with Event objects (such as retrieving and creating events) using interface
type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId string) (*Event, error)
	CreateOne(ctx context.Context, event Event) (*Event, error)
}

/* Each method accepts a context.Context as its first argument. 
This is used in Go for managing cancellation, timeouts, and other request-scoped values in long-running operations (like database queries). 
It's a way to pass metadata about a request through different layers of a program,
ensuring that operations can be cancelled or timed out if needed.
*/
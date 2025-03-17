package ports

import (
	"context"

	"github.com/google/uuid"
)

type Stage interface {
	// Unique identifier for the stage
	GetID() uuid.UUID

	// Main execution function
	Execute(ctx context.Context, input interface{}) (interface{}, error)

	// Error handling function
	HandleError(ctx context.Context, err error) error

	// Rollback function for failure recovery
	Rollback(ctx context.Context, input interface{}) error
}

package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// UserRepository defines methods the service layer expects
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}

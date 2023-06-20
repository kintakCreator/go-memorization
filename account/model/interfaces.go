package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

// TokenService defines methods the handler layer expects to interact
// with in regards to producing JWTs as string
type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// UserRepository defines methods the service layer expects
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}

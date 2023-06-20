package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jacobsgoodwin/memrizr/account/model"
	"github.com/jacobsgoodwin/memrizr/account/model/apperrors"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected into this
// service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// init a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}

// Signup reaches our to a UserRepository to verify the email addres
func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	//panic("Method not implemented")
	return apperrors.NewBadRequest("aboba")
}

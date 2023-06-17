package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jacobsgoodwin/memrizr/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// Get is mock of UserService Get
func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	// args that will be passed to "return" in the tests, when function
	// is called with a uid. Hence the name "ret"
	ret := m.Called(ctx, uid)

	// First value passed to "return"
	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

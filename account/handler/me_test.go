package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobsgoodwin/memrizr/account/model"
	"github.com/jacobsgoodwin/memrizr/account/model/apperrors"
	"github.com/jacobsgoodwin/memrizr/account/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "some@post.com",
			Name:  "Mike Nicklson",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// use a middleware to set context for test
		// the only claims the UID
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UID: uid,
			},
			)
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})

	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, 500, rr.Code)
		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down call chain"))

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UID: uid,
			},
			)
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})

		assert.Equal(t, respErr.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})
}

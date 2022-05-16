package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/service"
	mock_service "github.com/Korisss/skymp-master-api-go/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestHandler_register(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, user domain.User)

	tests := []struct {
		name                 string
		body                 string
		user                 domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			body: `{
				"name": "Koriss",
				"email": "koriss@koriss.com",
				"password": "123456789"
			}`,
			user: domain.User{
				Name:     "Koriss",
				Email:    "koriss@koriss.com",
				Password: "123456789",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user domain.User) {
				r.EXPECT().CreateUser(user).Return("1", nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":"1"}`,
		},
		{
			name: "Invalid request",
			body: `{
				"name": "Koriss",
				"password": "123456789"
			}`,
			user:                 domain.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user domain.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name: "Invalid email",
			body: `{
				"name": "Koriss",
				"email": "koriss",
				"password": "123456789"
			}`,
			user:                 domain.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user domain.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"mail: missing '@' or angle-addr"}`,
		},
		{
			name: "Invalid password",
			body: `{
				"name": "Koriss",
				"email": "koriss@koriss.com",
				"password": "1234"
			}`,
			user:                 domain.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user domain.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"password must contain at least 6 characters"}`,
		},
		{
			name: "Service Error",
			body: `{
				"name": "Koriss",
				"email": "koriss@koriss.com",
				"password": "123456789"
			}`,
			user: domain.User{
				Name:     "Koriss",
				Email:    "koriss@koriss.com",
				Password: "123456789",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user domain.User) {
				r.EXPECT().CreateUser(user).Return("0", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.user)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/users", handler.register)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/users",
				bytes.NewBufferString(test.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

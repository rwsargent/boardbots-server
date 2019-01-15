package newuser

import (
	"testing"
	"boardbots/server/mocks"
	"boardbots/server/testingutils"
	"net/http"
	"github.com/stretchr/testify/assert"
	"errors"
	"github.com/labstack/echo"
)

func TestNewUserSuccess(t *testing.T) {
  mockPortal := mocks.UserPortal{}
  mockPortal.On("NewUser", "username", "pass").Return(nil)
	handler := Handler{
		UserPortal: mockPortal,
	}

	ctx, rec := testingutils.FakeContext(http.MethodPost, "/newuser", `{"username" : "username", "password" : "pass"}`)

	if assert.NoError(t, handler.NewUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestNewUserFails(t *testing.T) {
	mockPortal := mocks.UserPortal{}
	mockPortal.On("NewUser", "username", "pass").Return(errors.New("could not compute"))
	handler := Handler{
		UserPortal: mockPortal,
	}
	ctx, _ := testingutils.FakeContext(http.MethodPost, "/newuser", `{"username" : "username", "password" : "pass"}`)
	result := handler.NewUser(ctx).(*echo.HTTPError)
	if assert.Error(t, result) {
		assert.Equal(t, http.StatusInternalServerError, result.Code)
	}
}
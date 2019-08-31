package newuser

import (
	"boardbots-server/server/mocks"
	"boardbots-server/server/testingutils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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

func TestNewUser_UserAlreadyExists(t *testing.T) {
	mockPortal := mocks.UserPortal{}
	mockPortal.On("NewUser", "username", "pass").Return(echo.NewHTTPError(http.StatusBadRequest, "username already exists"))

	handler := Handler{
		UserPortal: mockPortal,
	}
	ctx, _ := testingutils.FakeContext(http.MethodPost, "/newuser", `{"username" : "username", "password" : "pass"}`)
	result := handler.NewUser(ctx).(*echo.HTTPError)
	if assert.Error(t, result) {
		assert.Equal(t, http.StatusBadRequest, result.Code)
	}
}

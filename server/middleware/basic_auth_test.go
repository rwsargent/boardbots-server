package middleware

import (
	"testing"
	"boardbots/server/mocks"
	"github.com/stretchr/testify/assert"
	"boardbots/server/testingutils"
	"boardbots/server/context"
	"github.com/stretchr/testify/mock"
)

func TestBasicAuth(t *testing.T) {
	userPortal := mocks.UserPortal{}
	username := "test"
	password := "creds"
	userPortal.On("ValidateCredentials", username, password).Return(true, nil)
	userPortal.On("GetPlayerPrinciple", username).Return(context.PlayerPrinciple{}, nil)
	authenticator := GetBasicAuthenticator(userPortal)
	ctx, _ := testingutils.FakeBBContext()
	expected, err := authenticator.Validator(username, password, ctx)

	assert.NoError(t, err)
	assert.True(t, expected)
}

func TestAuthFails(t *testing.T) {
	userPortal := mocks.UserPortal{}
	userPortal.On("ValidateCredentials", mock.Anything, mock.Anything).Return(false, nil)
	// userPortal.On("GetPlayerPrinciple", username).Return(context.PlayerPrinciple{}, nil)

	authenticator := GetBasicAuthenticator(userPortal)
	ctx, _ := testingutils.FakeBBContext()
	expected, err := authenticator.Validator("", "", ctx)
	assert.Error(t, err)
	assert.False(t, expected)
}
package mocks

import (
	"github.com/stretchr/testify/mock"
	"boardbots/server/context"
)

type UserPortal struct {
	mock.Mock
}

func (portal UserPortal) ValidateCredentials(username, password string) bool {
	args := portal.Called(username, password)
	return args.Bool(0)
}

func (portal UserPortal) GetPlayerPrinciple(username string) (context.PlayerPrinciple, error) {
	args := portal.Called(username)
	return args.Get(0).(context.PlayerPrinciple), args.Error(1)
}

func (portal UserPortal) NewUser(username, password string) error {
	args := portal.Called(username, password)
	return args.Error(0)
}
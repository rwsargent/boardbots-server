package persistence

import (
	"boardbots/server/context"
	"errors"
	"github.com/google/uuid"
)

type (
	// Mocked
	UserPortal interface {
		ValidateCredentials(username, password string) bool
		GetPlayerPrinciple(username string) (context.PlayerPrinciple, error)
		NewUser(username, password string) error
	}

	InMemoryUsers struct {
		Credentials map[string]string
		Principles map[string]context.PlayerPrinciple
	}
)

func (portal *InMemoryUsers) ValidateCredentials(username, password string) bool {
	pass, ok := portal.Credentials[username]
	if !ok {
		return false
	}
	return pass == password
}

func (portal *InMemoryUsers) GetPlayerPrinciple(username string) (context.PlayerPrinciple, error) {
	if principle, ok := portal.Principles[username]; ok {
		return principle, nil
	} else {
		return context.PlayerPrinciple{}, errors.New("cannot find player")
	}
}

func (portal *InMemoryUsers) NewUser(username, password string) error {
	portal.Credentials[username] = password
	portal.Principles[username] = context.PlayerPrinciple{
		UserName: username,
		UserId : uuid.New(),
	}
	return nil
}

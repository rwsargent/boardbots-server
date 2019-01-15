package persistence

import (
	"boardbots/server/context"
	"errors"
)

type (
	UserPortal interface {
		ValidateCredentials(username, password string) bool
		GetPlayerPrinciple(username string) (context.PlayerPrinciple, error)
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

package persistence

import (
	"boardbots/server/context"
	"errors"
	"github.com/google/uuid"
	"boardbots/server/transport"
	"net/http"
	"sync"
)

type (
	// Mocked
	UserPortal interface {
		ValidateCredentials(username, password string) bool
		GetPlayerPrinciple(username string) (context.PlayerPrinciple, error)
		NewUser(username, password string) error
	}

	InMemoryUsers struct {
		lock sync.RWMutex
		Credentials map[string]string
		Principles map[string]context.PlayerPrinciple
	}
)
func NewMemoryPortal() *InMemoryUsers {
	return &InMemoryUsers{
		Credentials: make(map[string]string, 0),
		Principles: make(map[string]context.PlayerPrinciple, 0),
	}
}
func (portal *InMemoryUsers) ValidateCredentials(username, password string) bool {
	portal.lock.RLock()
	defer portal.lock.RUnlock()
	pass, ok := portal.Credentials[username]
	if !ok {
		return false
	}
	return pass == password
}

func (portal *InMemoryUsers) GetPlayerPrinciple(username string) (context.PlayerPrinciple, error) {
	portal.lock.RLock()
	defer portal.lock.RUnlock()
	if principle, ok := portal.Principles[username]; ok {
		return principle, nil
	} else {
		return context.PlayerPrinciple{}, errors.New("cannot find player")
	}
}

func (portal *InMemoryUsers) NewUser(username, password string) error {
	portal.lock.Lock()
	defer portal.lock.Unlock()
	if _, userExists :=  portal.Credentials[username]; userExists {
		return transport.NewHandledError(http.StatusConflict, "username already exists")
	}
	portal.Credentials[username] = password
	portal.Principles[username] = context.PlayerPrinciple{
		UserName: username,
		UserId : uuid.New(),
	}
	return nil
}

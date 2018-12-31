package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryGameManager_CreateGame(t *testing.T) {
	manager := InMemoryGameManager{}
	uuid := manager.Ad()

	assert.NotNil(t, uuid)
}

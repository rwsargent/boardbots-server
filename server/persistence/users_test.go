package persistence

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var portal InMemoryUsers

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	portal = InMemoryUsers{
		Credentials: make(map[string]string, 0),
	}
}
func TestUserDoesntExit(t *testing.T) {
	assert.False(t, portal.ValidateCredentials("user", "pass"))
}

func TestWhenUserExists_ValidateReturnsTrue(t *testing.T) {
	portal.Credentials["user"] = "pass"

	assert.True(t, portal.ValidateCredentials("user", "pass"))
}

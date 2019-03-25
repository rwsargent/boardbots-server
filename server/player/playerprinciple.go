package player

import "github.com/google/uuid"

type (
	PlayerPrinciple struct {
		UserName string
		UserId   uuid.UUID
		Password string
	}
)

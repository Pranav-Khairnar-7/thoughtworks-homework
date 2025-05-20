package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

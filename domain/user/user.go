package user

import (
	"time"
)

type User struct {
	Id        uint32
	Name      string
	Age       uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

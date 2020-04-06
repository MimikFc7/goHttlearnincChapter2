package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Age       int       `json:"age"`
	Birthdate int64     `json:"birthdate"`
}

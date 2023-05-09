package domain

import "github.com/google/uuid"

type Farmer struct {
	ID       uuid.UUID
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

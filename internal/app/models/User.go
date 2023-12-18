package models

import "time"

type User struct {
	ID             uint
	Username       string
	PasswordHash   string
	Email          string
	IsDeleted      bool
	RegisteredData time.Time
	CreateAt       time.Time
	UpdateAt       time.Time
}

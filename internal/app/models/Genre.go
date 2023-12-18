package models

import "time"

type Genre struct {
	ID        uint
	GenreName string
	IdDeleted bool
	CreateAt  time.Time
	UpdateAt  time.Time
}

package models

import "time"

type Author struct {
	ID         uint
	AuthorName string
	IsDeleted  bool
	CreateAt   time.Time
	UpdateAt   time.Time
}

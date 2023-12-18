package models

import "time"

type Collection struct {
	ID             uint
	CollectionName string
	IsDeleted      bool
	CreateAt       time.Time
	UpdateAt       time.Time
}

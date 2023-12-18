package models

import "time"

type Album struct {
	ID          uint
	AuthorID    uint
	AlbumName   string
	IsDeleted   bool
	ReleaseYear time.Time
	CreateAt    time.Time
	UpdateAt    time.Time
}

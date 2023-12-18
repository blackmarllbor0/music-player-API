package models

import "time"

type Track struct {
	ID          uint
	AuthorID    uint
	GenreID     uint
	AlbumID     uint
	IsDeleted   bool
	Duration    time.Duration
	ReleaseDate time.Time
	CreateAt    time.Time
	UpdateAt    time.Time
}

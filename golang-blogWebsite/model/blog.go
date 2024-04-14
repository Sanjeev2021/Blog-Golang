package model

import "time"

type Blog struct {
	ID        uint
	AuthorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Content   string
}

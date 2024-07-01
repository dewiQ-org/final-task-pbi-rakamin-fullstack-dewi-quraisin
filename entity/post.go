package entity

import "time"

type Post struct {
	ID        int
	UserID    int
	Picture   *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

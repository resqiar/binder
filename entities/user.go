package entities

import "time"

type User struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Provider   string
	Fullname   string
	Username   string
	Email      string
	Bio        string
	PictureURL string
}

package entities

import (
	"database/sql"
	"time"
)

type Extension struct {
	ID               string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Title            string
	Description      sql.NullString
	Code             sql.NullString
	YoutubeURL       sql.NullString
	AuthorID         string
	ImageAttachments []ImageAttachment
}

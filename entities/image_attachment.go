package entities

import "database/sql"

type ImageAttachment struct {
	ID          sql.NullString
	URL         sql.NullString
	ExtensionID sql.NullString
}

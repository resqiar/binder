package utils

import "context"

func GetUserIDFromContext(c context.Context) string {
	id, ok := c.Value("userID").(string)
	if !ok {
		return ""
	}
	return id
}

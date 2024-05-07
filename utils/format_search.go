package utils

import "strings"

func FormatSearch(keyword string) string {
	words := strings.Split(keyword, " ")
	parts := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			parts = append(parts, word)
		}
	}

	return strings.Join(parts, " & ")
}

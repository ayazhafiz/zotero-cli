package v3

import (
	"fmt"
	"strings"
)

type Author struct {
	FirstName string
	LastName  string
}

func PrintAuthors(authors []Author) string {
	if len(authors) == 0 {
		return "(unknown)"
	}

	lasts := make([]string, len(authors))
	for i, author := range authors {
		lasts[i] = author.LastName
	}

	if len(authors) <= 2 {
		return strings.Join(lasts, " and ")
	}
	if len(authors) == 3 {
		return fmt.Sprintf("%s, %s, and %s", lasts[0], lasts[1], lasts[2])
	}
	return fmt.Sprintf("%s et al.", lasts[0])
}

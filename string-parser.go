package stringparser

import (
	"strings"
)

func Parse(str string, seperator string) []string {

	items := strings.Fields(str)
	if seperator != "" {
		items = strings.Split(str, seperator)
	}

	return items
}

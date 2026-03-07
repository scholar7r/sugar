// Package strung provides tools to process strings
package strung

import "strings"

// TrimAround trims spaces around string
//
// Example:
//
// ```go
// v := TrimAround([]string{" ADD ", " 1 ", " AND ", " 2 "}) // v = []string{"ADD", "1", "AND", "2"}
// ```
func TrimAround(v []string, cutset string) []string {
	var trimmed []string

	for _, w := range v {
		trimmed = append(trimmed, strings.Trim(w, cutset))
	}

	return trimmed
}

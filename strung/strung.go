// Package strung provides tools to process strings
package strung

import "strings"

// TrimAround trims characters around string.
func TrimAround(v []string, cutset string) []string {
	var trimmed []string

	for _, w := range v {
		trimmed = append(trimmed, strings.Trim(w, cutset))
	}

	return trimmed
}

// Unique removes duplicated strings from a slice, preserving order.
func Unique(slice []string) []string {
	var (
		seen   = make(map[string]struct{}, len(slice))
		result = make([]string, 0, len(slice))
	)

	for _, v := range slice {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

package iteration

import "strings"

const repeatCount = 5

// repeat characters repeatCount times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatOptimized(character string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

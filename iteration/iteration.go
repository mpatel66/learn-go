package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder // instantiated as an empty string
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

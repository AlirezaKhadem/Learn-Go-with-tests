package iteration

import "strings"

func Repeat(character string) string {
	var stringBuilder strings.Builder
	stringBuilder.Grow(5)

	for index := 0; index < 5; index++ {
		stringBuilder.WriteString(character)
	}

	return stringBuilder.String()
}

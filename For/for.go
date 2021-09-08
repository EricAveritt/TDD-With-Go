package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, times int) string {
	return strings.Repeat(character, times)
}

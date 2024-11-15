package iteration

import "strings"

func Repeat(character string, times int) (repeated string) {
	// for i := 0; i < times; i++ {
	// 	repeated += character
	// }
	// repeated = strings.Repeat(character, times)
	var b strings.Builder
	for i := 0; i < times; i++ {
		b.WriteString(character)
	}
	repeated = b.String()
	return
}

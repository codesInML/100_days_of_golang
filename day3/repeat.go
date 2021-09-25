package iteration

// import (
// 	"fmt"
// )

func Repeat (character string, iter int) string {
	var repeated string

	for i := 0; i < iter; i++ {
		repeated += character
	}

	return repeated
}

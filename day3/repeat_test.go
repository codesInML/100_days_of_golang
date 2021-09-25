package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat (t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat () {
	repeated := Repeat("a", 8)
	fmt.Println(repeated)
	// Output: aaaaaaaa
}

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < 5; i++ {
// 		Repeat("a")
// 	}
// }
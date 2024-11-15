package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("iterate once", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		verifyExpected(t, repeated, expected)
	})
	t.Run("dont iterate", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		verifyExpected(t, repeated, expected)
	})
	t.Run("10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		verifyExpected(t, repeated, expected)
	})
}

func verifyExpected(t testing.TB, repeated, expected string) {
	if repeated != expected {
		t.Errorf("expedted %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

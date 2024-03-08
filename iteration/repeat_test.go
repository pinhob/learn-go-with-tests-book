package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 3)
	expected := "aaa"

	if repeat != expected {
		t.Errorf("expected %s but got %s", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}

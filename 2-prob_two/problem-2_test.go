package prob_two

import (
	"testing"
)

func TestFibEvenSum(t *testing.T) {
	want := 4613732

	got := FibEvenSum(4_000_000)
	if want != got {
		t.Errorf("Wanted: %d, got: %d", want, got)
	}

}

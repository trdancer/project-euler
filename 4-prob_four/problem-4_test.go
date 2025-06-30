package prob_four_test

import (
	"testing"

	prob_four "github.com/trdancer/projeuler/4-prob_four"
)

func TestProblemFour(t *testing.T) {
	want := 906609
	got := prob_four.LargestPalindrome(3)
	if want != got {
		t.Errorf("Expected: %d, got: %d", want, got)
	}
}

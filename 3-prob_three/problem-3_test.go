package prob_three_test

import (
	"testing"

	prob_three "github.com/trdancer/projeuler/3-prob_three"
)

func TestProblemThree(t *testing.T) {
	want := 6857
	got := prob_three.LargestPrimeFactor(600851475143)
	if got != want {
		t.Errorf("Expected: %d, got: %d", want, got)
	}
}

func TestGetFactor(t *testing.T) {
	want1, want2 := prob_three.GetLargestFactor(13)
	got1, got2 := 13, true
	if want1 != got1 {
		t.Errorf("Did not return largest factor. Expected %d, got %d", want1, got1)
	}
	if want2 != got2 {
		t.Errorf("Did not return isPrime correctly. Expected %t, got %t", want2, got2)
	}
}

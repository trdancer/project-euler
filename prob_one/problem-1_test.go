package prob_one

import t "testing"

func TestProblemOne(t *t.T) {
	want := 233168
	got := Multiples(1000)

	if got != want {
		t.Error(want, got)
	}

}

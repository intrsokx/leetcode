package base_test

import "testing"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := Abs(-1)

	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

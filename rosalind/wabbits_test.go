package rosalind

import "testing"

func TestWabbits(t *testing.T) {
	rslt := Wabbits(5, 3)
	if rslt != 19 {
		t.Errorf("Wabbits(5, 3) == %v, expected 19", rslt)
	}
}

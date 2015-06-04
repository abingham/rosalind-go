package rosalind

import "testing"

func equal(a []byte, b []byte) (rslt bool) {
	rslt = false;
	
	if len(a) != len(b) {
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return
		}
	}

	rslt = true
	return
}

func TestReverseComplement(t *testing.T) {
	input :=    []byte("AAAACCCGGT")
	expected := []byte("ACCGGGTTTT")
	rslt := ReverseComplement(input)
	if !equal(expected, rslt) {
		t.Errorf("ReverseComplement(%q) == %q, want %q",
			input, rslt, expected)
	}
}

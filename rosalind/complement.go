package rosalind

// Get the complement of a single DNA base
func baseComp(base byte) (comp byte) {
	switch base {
	case 'A':
		comp = 'T'
	case 'T':
		comp = 'A'
	case 'C':
		comp = 'G'
	case 'G':
		comp = 'C'
	}
	return
}

// Get the reverse complement of a DNA string
func ReverseComplement(dat []byte) (rcomp []byte) {
	rcomp = make([]byte, len(dat))
	n := len(dat)
	for i := 0; i < n / 2; i++ {
		rcomp[i], rcomp[n - 1 - i] = baseComp(dat[n - 1 - i]), baseComp(dat[i])
	}
	return
}


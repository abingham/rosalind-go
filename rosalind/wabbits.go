package rosalind

func Wabbits(n int, k int) (curr int) {
	prev := 1
	curr = 1

	for i := 0; i < n-2; i++ {
		curr, prev = prev*k+curr, curr
	}

	return
}

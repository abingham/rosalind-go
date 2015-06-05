package rosalind

func Wabbits(n int, k int) (curr int) {
	prev := 0
	curr = 1

	for i := 0; i < n; i++ {
		prev, curr =  curr, (prev + curr) * k
	}
	
	return
}

package rosalind

import "sort"
import "github.com/cznic/sortutil"

func CountChars(dat string) (counts map[rune]int) {
	counts = make(map[rune]int)
	for _, char := range dat {
		counts[char] = counts[char] + 1
	}
	return
}

func SortKeys(counts map[rune]int) (keys []rune) {
	keys = make([]rune, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Sort(sortutil.RuneSlice(keys))
	return
}

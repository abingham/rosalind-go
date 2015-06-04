package main

import "fmt"
import "os"
import "sort"
import "strings"

import "github.com/abingham/rosalind-go/util"
import "github.com/cznic/sortutil"

func countChars(dat string) (counts map[rune]int) {
	counts = make(map[rune]int)
	for _, char := range dat {
		counts[char] = counts[char] + 1
	}
	return
}

func sortKeys(counts map[rune]int) (keys []rune) {
	keys = make([]rune, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Sort(sortutil.RuneSlice(keys))
	return
}

func main() {
	filename := os.Args[1]
	dat := strings.TrimSpace(util.ReadFileString(filename))
	counts := countChars(dat)
	keys := sortKeys(counts)
	
	for _, key := range keys {
		fmt.Print(counts[key], " ")
	}
	fmt.Print("\n")
}

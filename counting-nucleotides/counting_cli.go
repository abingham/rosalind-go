package main

import "fmt"
import "os"
import "strings"

import "github.com/abingham/rosalind-go/rosalind"

func main() {
	filename := os.Args[1]
	dat := strings.TrimSpace(rosalind.ReadFileString(filename))
	counts := rosalind.CountChars(dat)
	keys := rosalind.SortKeys(counts)
	
	for _, key := range keys {
		fmt.Print(counts[key], " ")
	}
	fmt.Print("\n")
}

package main

import "fmt"
import "os"
import "strings"

import "github.com/abingham/rosalind-go/rosalind"

func main() {
	filename := os.Args[1]
	dat := []byte(strings.TrimSpace(rosalind.ReadFileString(filename)))
	fmt.Println(string(rosalind.ReverseComplement(dat)))
}

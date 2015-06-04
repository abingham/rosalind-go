package main

import "os"
import "strings"
import "fmt"

import "github.com/abingham/rosalind-go/rosalind"

func main() {
	filename := os.Args[1]
	dat := strings.TrimSpace(rosalind.ReadFileString(filename))
	dat = strings.Replace(dat, "T", "U", -1)
	fmt.Println(dat)
}

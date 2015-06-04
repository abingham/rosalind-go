package main

import "os"
import "strings"
import "fmt"
import "github.com/abingham/rosalind-go/util"

func main() {
	filename := os.Args[1]
	dat := strings.TrimSpace(util.ReadFileString(filename))
	dat = strings.Replace(dat, "T", "U", -1)
	fmt.Println(dat)
}

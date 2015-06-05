package main

import "fmt"
import "os"
import "strconv"

import "github.com/abingham/rosalind-go/rosalind"

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 32)
	rosalind.Check(err)
	k, err := strconv.ParseInt(os.Args[2], 10, 32)
	rosalind.Check(err)
	
	fmt.Println(rosalind.Wabbits(int(n), int(k)))
}

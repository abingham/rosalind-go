package main

import "fmt"
import "io/ioutil"
import "os"
import "sort"
import "strings"

import "github.com/cznic/sortutil"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	filename := os.Args[1]
	byte_dat, err := ioutil.ReadFile(filename)
	check(err)

	dat := string(byte_dat)
	dat = strings.TrimSpace(dat)
	
	counts := make(map[rune]int)
	
	for _, char := range dat {
		counts[char] = counts[char] + 1
	}

	var keys sortutil.RuneSlice = make([]rune, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Sort(keys)
	
	for _, key := range keys {
		fmt.Print(counts[key], " ")
	}
	fmt.Print("\n")
}

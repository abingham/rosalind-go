package main

import "io/ioutil"
import "os"
import "strings"
import "fmt"

// TODO: into lib
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// TODO: Put this in a lib of some sort
func readFileString(filename string) (dat string) {
	byte_dat, err := ioutil.ReadFile(filename)
	check(err)
	dat = string(byte_dat)
	return
}

func main() {
	filename := os.Args[1]
	dat := strings.TrimSpace(readFileString(filename))
	dat = strings.Replace(dat, "T", "U", -1)
	fmt.Println(dat)
}

// Various common functions for the rosalind-go project
package util

import (
	"io/ioutil"
)

// Check if there has been an error, panicking of so.
func Check(e error) {
    if e != nil {
        panic(e)
    }
}

// Read the entire contents of a file as a string.
// The data is not trimmed.
func ReadFileString(filename string) (dat string) {
	byte_dat, err := ioutil.ReadFile(filename)
	Check(err)
	dat = string(byte_dat)
	return
}

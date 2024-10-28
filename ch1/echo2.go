package main

import (
	"fmt"
	"os"
)

//implementation of Unix echo command
// "go run echo2.go ... ..."

func main () {
	s, sep := "", ""
	// use blank identifier when syntax requires a variable name but program logic does not.
	// range returns both index and value of element at index
	for _, arg := range os.Args[1:] { // _ discards unwanted loop index, returns only arg
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// s := ""
// var s string
// var s = ""
// var s string = "”

// use one of first two
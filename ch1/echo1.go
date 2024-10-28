package main

import (
	"fmt"
	"os"
)

//implementation of Unix echo command

func main() {
	var s, sep string //declares 2 variables: s and sep of type string
	// := is short variable declaration; 
	for i := 1; i < len(os.Args); i++ { //go also has range
		s += sep + os.Args[i] // + concatenates strings, += is assignment operator
		sep = " " //sep initially empty, set to a space after first iteration
	}
	fmt.Println(s)
}

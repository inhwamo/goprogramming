package main

import (
	"fmt"
	"os"
)

//implementation of Unix echo command

func main() {
	var s, sep string //declares 2 variables: s and sep of type string
	for i := 1; i < len(os.Args); i++ { //go also has range
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

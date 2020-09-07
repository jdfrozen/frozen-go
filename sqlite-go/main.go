package main

import (
	"fmt"
)

func main() {
	initCache()
	for {
		var input string
		fmt.Scanln(&input)
		vm(input)
		fmt.Printf("> ")
	}

}

package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Scanln(&input)
		vm(input)
		fmt.Printf("> ")
	}

}

package main

import (
	"fmt"
)

func main() {
	frozen := Row{1, "frozen", 2}
	var rowbs = createBytes(frozen)
	writeDb(rowbs)
	var rowbsr = readDb(35)
	var row = createRow(rowbsr)
	fmt.Println(row)
	for {
		var input string
		fmt.Scanln(&input)
		vm(input)
		fmt.Printf("> ")
	}

}

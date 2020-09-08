package main

import (
	"fmt"
)

func main() {
	id := 8
	var a = [10]int{1, 2, 3, 4, 5, 7, 7, 8, 9, 19}
	var startIndex = 0
	var endIndex = len(a) - 1
	var index int
	for endIndex > startIndex {
		index = (startIndex + endIndex) / 2
		v := a[index]
		if v == id {
			break
		}
		if v > id {
			endIndex = index
		} else {
			startIndex = index
		}
		if (endIndex - startIndex) == 1 {
			index = endIndex
			break
		}
	}
	fmt.Println(index)
	for {
		var input string
		fmt.Scanln(&input)
		vm(input)
		fmt.Printf("> ")
	}

}

package main

import (
	"fmt"
)

func main() {
	id := 3
	var a = [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	var startIndex = 0
	var endIndex = len(a) - 1
	var index int
	for endIndex >= startIndex {
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
		//在中间
		if (endIndex - startIndex) <= 1 {
			if id > a[endIndex] {
				index = endIndex + 1
				break
			}
			if id < a[startIndex] {
				index = startIndex
				break
			}
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

package main

import (
	"fmt"
)

func main() {
	//path := "frozen.db2"
	//f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	//defer f.Close()
	//if err != nil {
	//	panic(err)
	//}
	//n,_:=f.Seek(0,0)
	//fmt.Println(n)
	//f.WriteAt([]byte("55"), 1)

	for {
		var input string
		fmt.Scanln(&input)
		vm(input)
		fmt.Printf("> ")
	}

}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func createBytes(row Row) {
	//id
	var id = row.id
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, id)
	var idb = buf.Bytes()

	//age
	var age = row.age
	bufAge := bytes.NewBuffer(make([]byte, 0))
	binary.Write(bufAge, binary.BigEndian, age)
	var ageb = bufAge.Bytes()

	var nameb = row.name
	var rowb = [35]byte{}
	var idbLen = len(idb)
	var namebLen = len(nameb)
	for i, b := range idb {
		rowb[i] = b
	}
	for i, b := range nameb {
		rowb[i+idbLen] = b
	}
	for i, b := range ageb {
		rowb[i+namebLen+idbLen] = b
	}
	fmt.Println(rowb)
}

func main() {

	var str = "frozen"
	var datas = []byte(str)
	var nameb = [32]byte{}
	for i, data := range datas {
		nameb[i] = data
	}
	frozen := Row{1, nameb, 28}
	createBytes(frozen)
	for {
		var input string
		fmt.Scanf("%s", &input)
		fmt.Printf("> ")
	}

}

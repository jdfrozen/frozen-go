package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Row struct {
	id   uint16
	name [32]byte
	age  uint8
}

func createBytes(row Row) [35]byte {
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
	return rowb
}

func createRow(rowbs [35]byte) Row {
	var idb = make([]byte, 2)
	var nameb [32]byte
	var ageb = make([]byte, 1)
	var idbLen = len(idb)
	var namebLen = len(nameb)
	for i, _ := range idb {
		idb[i] = rowbs[i]
	}
	for i, _ := range nameb {
		nameb[i] = rowbs[i+idbLen]
	}
	for i, _ := range ageb {
		ageb[i] = rowbs[i+namebLen+idbLen]
	}
	buf := bytes.NewBuffer(idb)
	var id uint16
	binary.Read(buf, binary.BigEndian, &id)
	agebuf := bytes.NewBuffer(ageb)
	var age uint8
	binary.Read(agebuf, binary.BigEndian, &age)
	row := Row{id, nameb, age}
	return row
}

func main() {

	var str = "frozen"
	var datas = []byte(str)
	var nameb = [32]byte{}
	for i, data := range datas {
		nameb[i] = data
	}
	frozen := Row{1, nameb, 2}
	var rowbs = createBytes(frozen)
	var row = createRow(rowbs)
	fmt.Println(row.id)
	fmt.Println(row.age)
	fmt.Println(string(row.name[:]))
}

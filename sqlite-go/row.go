package main

import (
	"bytes"
	"encoding/binary"
)

type Row struct {
	id   uint16
	name string
	age  uint8
}

func createBytes(row Row) []byte {
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
	//序列化id
	var rowb = make([]byte, 35)
	var idbLen = len(idb)
	for i, b := range idb {
		rowb[i] = b
	}
	//序列化name
	var name = row.name
	var datas = []byte(name)
	var nameb = [32]byte{}
	for i, data := range datas {
		nameb[i] = data
	}
	for i, b := range nameb {
		rowb[i+idbLen] = b
	}
	//序列化age
	var namebLen = len(nameb)
	for i, b := range ageb {
		rowb[i+namebLen+idbLen] = b
	}
	return rowb
}

func createRow(rowbs []byte) Row {
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
	row := Row{id, string(nameb[:]), age}
	return row
}

package main

import "fmt"

var buffer = make([]byte, 4096)
var maxRow = (4096 / 35)
var index = 0

func insert(row Row) {
	var indexRow = (index + 1) / 35
	if indexRow >= maxRow {
		panic("保存超过最大行数")
	}
	var bytes = createBytes(row)
	for _, b := range bytes {
		buffer[index] = b
		index++
	}
}

func selectAll() {
	var i, a = 0, 0
	var bytes = make([]byte, 35)
	for i = 0; i < index; i++ {
		bytes[a] = buffer[i]
		if a == 34 {
			var row = createRow(bytes)
			fmt.Println(row)
			a = 0
		} else {
			a++
		}
	}
}

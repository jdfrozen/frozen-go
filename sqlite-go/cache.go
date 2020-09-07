package main

import "fmt"

var pager = Pager{0, 0, make([]Row, 116)}

func insert(row Row) {
	var indexRow = pager.rowNum
	if indexRow >= 116 {
		panic("保存超过最大行数")
	}
	pager.rows[pager.rowNum] = row
	pager.rowNum++
}

func selectAll() {
	var i uint16 = 0
	for ; i < pager.rowNum; i++ {
		fmt.Println(pager.rows[i])
	}
}

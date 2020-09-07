package main

import "fmt"

var pager = Pager{0, 0, make([]Row, 116)}

func insert(row Row) {
	var indexRow = pager.rowNum
	if indexRow >= 116 {
		panic("保存超过最大行数")
	}
	pager.rowNum++
	pager.rows[pager.rowNum] = row
}

func selectAll() {
	for _, row := range pager.rows {
		fmt.Println(row)
	}
}

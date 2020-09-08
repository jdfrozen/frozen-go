package main

import "fmt"

var pager = Pager{0, 0, make([]Row, 116)}

func insert(row Row) {
	var indexRow = pager.rowNum
	if indexRow >= 116 {
		panic("保存超过最大行数")
	}
	//为空直接插入
	if pager.rowNum == 0 {
		pager.rows[pager.rowNum] = row
		pager.rowNum++
	}
	//寻找插入的位置（二分查找）

}

func selectAll() {
	var i uint16 = 0
	for ; i < pager.rowNum; i++ {
		fmt.Println(pager.rows[i])
	}
}

//折半查找
func getIndex(rows []Row, id uint16) int {
	var startIndex = 0
	var endIndex = len(rows) - 1
	startRow := rows[startIndex]
	endRow := rows[endIndex]
	var index int
	for startRow.id >= endRow.id {
		index = (startIndex + endIndex + 1) / 2
		v := rows[index]
		if v.id == id {
			break
		}
		if v.id > id {
			endIndex = index
		} else {
			startIndex = index
		}
	}
	return index
}

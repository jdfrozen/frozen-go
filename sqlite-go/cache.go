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
		return
	}
	//寻找插入的位置（二分查找）
	index := getIndex(row.id)
	//数组数据copy移动
	rows := pager.rows
	var i = int(pager.rowNum)
	for ; i > index; i-- {
		rows[i] = rows[i-1]
	}
	rows[index] = row
	pager.rowNum++
}

func selectAll() {
	var i uint16 = 0
	for ; i < pager.rowNum; i++ {
		fmt.Println(pager.rows[i])
	}
}

//折半查找
func getIndex(id uint16) int {
	rows := pager.rows
	var startIndex = 0
	var endIndex = int(pager.rowNum) - 1
	var index = endIndex
	for endIndex >= startIndex {
		index = (startIndex + endIndex) / 2
		v := rows[index].id
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
			if id > rows[endIndex].id {
				index = endIndex + 1
				break
			}
			if id < rows[startIndex].id {
				index = startIndex
				break
			}
			index = endIndex
			break
		}
	}
	return index
}

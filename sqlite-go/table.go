package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Table struct {
	rowNum uint64
}

func initTable(table Table) {
	//rowNumindex
	var rowNum = table.rowNum
	bufAge := bytes.NewBuffer(make([]byte, 0))
	binary.Write(bufAge, binary.BigEndian, rowNum)
	var rowNumb = bufAge.Bytes()
	//序列化
	var tableb = make([]byte, 4096)
	//序列化rowNum
	for i, b := range rowNumb {
		tableb[i] = b
	}
	writeDb(0, tableb)
}

func readTable() Table {
	var table = Table{0}
	tablePage := readPagerOne(0)
	var rootPageIndexb = make([]byte, 8)
	var rowNumb = make([]byte, 8)
	var indexbLen = len(rootPageIndexb)
	for i, _ := range rootPageIndexb {
		rootPageIndexb[i] = tablePage[i]
	}
	for i, _ := range rowNumb {
		rowNumb[i] = tablePage[i+indexbLen]
	}
	agebuf := bytes.NewBuffer(rowNumb)
	var rowNum uint64
	binary.Read(agebuf, binary.BigEndian, &rowNum)
	table.rowNum = rowNum
	rootPage := readPagerOne(1)
	pageNew := readerPager(rootPage)
	var i uint16 = 0
	for ; i < pageNew.rowNum; i++ {
		fmt.Println(pageNew.rows[i])
	}
	return table
}

func insert(row Row) error {
	table := readTable()
	var rootPage []byte
	var pager Pager
	if table.rowNum <= 0 {
		pager = Pager{1, 0, make([]Row, 3)}
	} else {
		index := table.rowNum/116 + 1
		rootPage = readPagerOne(int64(index))
		pager = readerPager(rootPage)
	}
	table.rowNum++
	initTable(table)
	var indexRow = pager.rowNum
	if indexRow >= 116 {
		panic("保存超过最大行数")
	}
	//为空直接插入
	if pager.rowNum == 0 {
		pager.rows[pager.rowNum] = row
		pager.rowNum++
		writeDb(int64(pager.index), savePager(pager))
		return nil
	}
	//寻找插入的位置（二分查找）
	index, err := getIndex(row.id, pager)
	if err != nil {
		return err
	}
	//数组数据copy移动
	rows := pager.rows
	var i = int(pager.rowNum)
	for ; i > index; i-- {
		rows[i] = rows[i-1]
	}
	rows[index] = row
	pager.rowNum++
	writeDb(int64(pager.index), savePager(pager))
	return nil
}

//折半查找
func getIndex(id uint16, pager Pager) (int, error) {
	rows := pager.rows
	var startIndex = 0
	var endIndex = int(pager.rowNum) - 1
	var index = endIndex
	for endIndex >= startIndex {
		index = (startIndex + endIndex) / 2
		v := rows[index].id
		if v == id {
			return -1, errors.New("row exist !")
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
	return index, nil
}

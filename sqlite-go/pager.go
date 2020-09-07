package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Pager struct {
	index  uint16
	rowNum uint16
	rows   []Row
}

func savePager(pager Pager) []byte {
	//id
	var index = pager.index
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, index)
	var indexb = buf.Bytes()
	//age
	var rowNum = pager.rowNum
	bufAge := bytes.NewBuffer(make([]byte, 0))
	binary.Write(bufAge, binary.BigEndian, rowNum)
	var rowNumb = bufAge.Bytes()
	//序列化id
	var pageb = make([]byte, 4)
	var indexbLen = len(indexb)
	for i, b := range indexb {
		pageb[i] = b
	}
	//序列化age
	for i, b := range rowNumb {
		pageb[i+indexbLen] = b
	}
	return pageb
}

func readerPager(onePage []byte) Pager {
	var indexb = make([]byte, 2)
	var rowNumb = make([]byte, 2)
	var indexbLen = len(indexb)
	for i, _ := range indexb {
		indexb[i] = onePage[i]
	}
	for i, _ := range rowNumb {
		rowNumb[i] = onePage[i+indexbLen]
	}
	buf := bytes.NewBuffer(indexb)
	var index uint16
	binary.Read(buf, binary.BigEndian, &index)
	agebuf := bytes.NewBuffer(rowNumb)
	var rowNum uint16
	binary.Read(agebuf, binary.BigEndian, &rowNum)
	if rowNum == 0 {
		return Pager{index, rowNum, nil}
	}
	var rowIndex = 4
	maxRowIndex := int(rowNum * 35)
	for rowIndex <= maxRowIndex {
		if rowIndex == 34 {
			var row = createRow(onePage)
			fmt.Println(row)
			rowIndex = 0
		} else {
			rowIndex++
		}
	}
	var rows = make([]Row, 0)
	pager := Pager{index, rowNum, rows}
	return pager
}

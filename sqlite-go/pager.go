package main

import (
	"bytes"
	"encoding/binary"
)

type Pager struct {
	index  uint16
	rowNum uint16
	rows   []Row
}

func savePager(pager Pager) []byte {
	//index
	var index = pager.index
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, index)
	var indexb = buf.Bytes()
	//rowNumindex
	var rowNum = pager.rowNum
	bufAge := bytes.NewBuffer(make([]byte, 0))
	binary.Write(bufAge, binary.BigEndian, rowNum)
	var rowNumb = bufAge.Bytes()
	//序列化
	var pageb = make([]byte, 4096)
	var indexbLen = len(indexb)
	for i, b := range indexb {
		pageb[i] = b
	}
	//序列化rowNum
	for i, b := range rowNumb {
		pageb[i+indexbLen] = b
	}
	//序列化rows
	var i uint16 = 0
	var pagebIndex = 4
	for ; i < pager.rowNum; i++ {
		row := pager.rows[i]
		rowb := createBytes(row)
		for _, b := range rowb {
			pageb[pagebIndex] = b
			pagebIndex++
		}
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
	var prIndex = 4
	var rowIndex = 0
	maxPrIndex := int(rowNum * 35)
	var rows = make([]Row, rowNum)
	for prIndex <= maxPrIndex {
		var rowbs = make([]byte, 35)
		for i, _ := range rowbs {
			rowbs[i] = onePage[prIndex]
			prIndex++
		}
		var row = createRow(rowbs)
		rows[rowIndex] = row
		rowIndex++
	}

	pager := Pager{index, rowNum, rows}
	return pager
}

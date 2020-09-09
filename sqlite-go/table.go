package main

import (
	"bytes"
	"encoding/binary"
)

type Table struct {
	rootPageIndex uint64
	rowNum        uint64
}

func readTable() Table {
	onePage := readPagerOne(1)
	var rootPageIndexb = make([]byte, 8)
	var rowNumb = make([]byte, 8)
	var indexbLen = len(rootPageIndexb)
	for i, _ := range rootPageIndexb {
		rootPageIndexb[i] = onePage[i]
	}
	for i, _ := range rowNumb {
		rowNumb[i] = onePage[i+indexbLen]
	}
	buf := bytes.NewBuffer(rootPageIndexb)
	var rootPageIndex uint64
	binary.Read(buf, binary.BigEndian, &rootPageIndex)
	agebuf := bytes.NewBuffer(rowNumb)
	var rowNum uint64
	binary.Read(agebuf, binary.BigEndian, &rowNum)
	table := Table{rootPageIndex, rowNum}
	return table
}

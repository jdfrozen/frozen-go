package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Pager struct {
	index  uint16
	rowNum uint16
	rows   []Row
}

func savePager(pager Pager) {
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
	var pagerb = make([]byte, 4096)
	var indexbLen = len(indexb)
	for i, b := range indexb {
		pagerb[i] = b
	}
	//序列化rowNum
	for i, b := range rowNumb {
		pagerb[i+indexbLen] = b
	}
	//序列化rows
	var i uint16 = 0
	var pagebIndex = 4
	for ; i < pager.rowNum; i++ {
		row := pager.rows[i]
		rowb := createBytes(row)
		for _, b := range rowb {
			pagerb[pagebIndex] = b
			pagebIndex++
		}
	}
	writeDb(int64(pager.index), pagerb)
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

var pager = Pager{0, 0, make([]Row, 116)}

func insert(row Row) error {
	var indexRow = pager.rowNum
	if indexRow >= 116 {
		panic("保存超过最大行数")
	}
	//为空直接插入
	if pager.rowNum == 0 {
		pager.rows[pager.rowNum] = row
		pager.rowNum++
		return nil
	}
	//寻找插入的位置（二分查找）
	index, err := getIndex(row.id)
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
	return nil
}

func selectAll() {
	var i uint16 = 0
	for ; i < pager.rowNum; i++ {
		fmt.Println(pager.rows[i])
	}
}

//折半查找
func getIndex(id uint16) (int, error) {
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

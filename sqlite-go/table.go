package main

type Row struct {
	id   uint16
	name [32]byte
	age  uint8
}

type Pager struct {
	index   uint16
	rowNums uint16
}

type table struct {
	index   uint16
	pager   *Pager
	rowNums uint16
}

//游标定义
type Cursor struct {
	Table      *table
	pageNum    uint16
	cellNum    uint16
	endOfTable bool
}

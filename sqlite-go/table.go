package main

type table struct {
	root   uint64
	rowNum uint16
}

//游标定义
type Cursor struct {
	Table      *table
	pageNum    uint16
	cellNum    uint16
	endOfTable bool
}

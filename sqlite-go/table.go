package main

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

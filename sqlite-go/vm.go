package main

import (
	"fmt"
	"strconv"
	"strings"
)

func vm(str string) {
	if strings.Contains(str, "insert") {
		arr := strings.Split(str, ",")
		if len(arr) != 4 {
			inputErr()
			return
		}
		if arr[0] != "insert" {
			inputErr()
			return
		}
		id, r := strconv.Atoi(arr[1])
		if nil != r {
			inputErr()
			return
		}
		age, r := strconv.Atoi(arr[3])
		if nil != r {
			inputErr()
			return
		}
		row := Row{uint16(id), arr[2], uint8(age)}
		err := insert(row)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if strings.Contains(str, "select") {
		readTable()
		return
	}
	if strings.Contains(str, "init") {
		initTable(Table{0})
		return
	}
	inputErr()
}

func inputErr() {
	fmt.Println("Please enter the correct command !")
}

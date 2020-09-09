package main

import (
	"fmt"
	"os"
)

func checkFileIsExist() bool {
	path := "frozen.db"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeDb(index int64, bytes []byte) {
	path := "frozen.db"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	off := index * 4096
	length, err := f.WriteAt(bytes, off)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
}

func readPagerOne(index int64) []byte {
	if checkFileIsExist() {
		return make([]byte, 4096)
	}
	file, err := os.Open("frozen.db")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//从文件中读取len个字节
	bytes := make([]byte, 4096)
	off := index * 4096
	length, err := file.ReadAt(bytes, off)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("number of bytes read: %d\n", length)
	return bytes
}

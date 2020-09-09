package main

import (
	"fmt"
	"os"
)

func writeDb(index int64, bytes []byte) {
	path := "frozen.db"
	f, err := os.Create(path)
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

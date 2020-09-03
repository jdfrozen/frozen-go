package main

import (
	"fmt"
	"log"
	"os"
)

func writeDb(bytes []byte) {
	path := "frozen.db"
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	length, err := f.WriteAt(bytes, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
}

func readDb(len int64) []byte {
	file, err := os.Open("frozen.db")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//从文件中读取16个字节
	bytes := make([]byte, len)
	length, err := file.ReadAt(bytes, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("number of bytes read: %d\n", length)
	return bytes
}

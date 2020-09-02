package main

import (
	"log"
	"os"
)

func write() {
	file, err := os.Open("frozen.db")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//从文件中读取16个字节
	bytes := make([]byte, 16)
	br, err := file.WriteAt(bytes, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("number of bytes write: %d\n", br)
	log.Printf("Data read: %s\n", bytes)
}

func read() {
	file, err := os.Open("frozen.db")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//从文件中读取16个字节
	bytes := make([]byte, 16)
	br, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("number of bytes read: %d\n", br)
	log.Printf("Data read: %s\n", bytes)
}

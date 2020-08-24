package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	args := os.Args
	names := args[1:]
	// 打开文件
	f, err := os.Open(names[0])
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)
	go func() {
		for {
			line, err := rd.ReadBytes('\n')
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
				continue
			} else if err != nil {
				panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
			}
			s := string(line[:len(line)-1])
			fmt.Println(s)
		}
	}()
	time.Sleep(500000 * time.Millisecond)
}

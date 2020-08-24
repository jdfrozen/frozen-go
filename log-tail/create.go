package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var Info *log.Logger

func init() {
	infoFile, err := os.OpenFile("infos.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		age := rand.Intn(3)
		Info.Println("frozen age:", age)
		//等一秒钟
		time.Sleep(1000 * time.Millisecond)
	}
}

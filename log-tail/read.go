package main

import (
	"bufio"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"io"
	"log"
	"os"
	"time"
)

func writeInfluxDb(info chan string) {
	for {
		// https://github.com/influxdata/influxdb/tree/master/client
		infClient, err := client.NewHTTPClient(client.HTTPConfig{
			Addr:     "http://192.168.226.130:8086",
			Username: "admin",
			Password: "",
		})
		if err != nil {
			panic(fmt.Sprintf("influxdb NewHTTPClient err:%s", err.Error()))
		}
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			//数据库
			Database: "my_test",
			//设置提供的Unix时间值的精度。InfluxDB假设您没有指定时间戳以纳秒为单位precision
			Precision: "s",
		})
		if err != nil {
			panic(fmt.Sprintf("influxdb NewBatchPoints error:%s", err.Error()))
		}
		//组装数据，tags 作为查询条件
		tags := map[string]string{"name": "info"}
		//组装具体值
		value := <-info
		fmt.Println(value)
		//组装具体值
		fields := map[string]interface{}{
			"id":   1,
			"sex":  1,
			"pass": 0707,
		}
		//表——myuser
		pt, err := client.NewPoint("myuser", tags, fields, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
		//写入
		err = infClient.Write(bp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	// 打开文件
	f, err := os.Open("infos.log")
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)
	info := make(chan string)
	go writeInfluxDb(info)
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
			info <- s
		}
	}()
	time.Sleep(500000 * time.Millisecond)
}

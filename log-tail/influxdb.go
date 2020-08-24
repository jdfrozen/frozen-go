package main

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

func main() {
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
	tags := map[string]string{"name": "xc"}
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

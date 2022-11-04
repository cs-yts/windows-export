package model

import (
	"fmt"
	"time"
	"windows_export/utils"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)


//influxDB 全局链接
var InfluxClient influxdb2.Client

//influxDB初始化
func InfluxInit() {
	//创建客户端
	InfluxClient = influxdb2.NewClient(utils.InfluxDBUrl, utils.InfluxdbToken)
}

//关闭influxdb
func InfluxClose() {
	InfluxClient.Close()
}

//数据插入行协议组成
// |measurement|tag_set|field_set|timestap|
// field_set 可以是整数,浮点数,字符串和布尔值,默认是以浮点数来存储,如果要以整数来存储,则在数值后面加i 例如:35i ,布尔值:t,T,True,true/f,F,false,False,其它2个就是默认类型
// 特殊字符转义都是通过\来实现: , \, , = \=  , 空格 也是通过 \ 来转义,
//写入influxDB
func InfluxWrite(measurement, tag, ip, field string, value interface{}) {
	//连接端口
	writeAPI := InfluxClient.WriteAPI(utils.InfluxOrg, utils.InfluxBucket)
	//捕获错误
	errorsCh := writeAPI.Errors()
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()
	//NewPoint创建一个模板
	p := influxdb2.NewPoint(measurement, //
		map[string]string{
			ip: ""},
		map[string]interface{}{
			field: value},
		time.Now())

	// NewPointWithMeasurement 根据上面创建的模板填充数据
	p1 := influxdb2.NewPointWithMeasurement(measurement).
		AddTag(tag, ip).
		AddField(field, value).
		SetTime(time.Now())

	//write point asynchronously
	writeAPI.WritePoint(p)
	writeAPI.WritePoint(p1)
	// Flush writes
	writeAPI.Flush()
}

// middleware.InfluxWrite("System", "system_load", "172.16.7.10", "system_load", loadInfo)

//查询数据
// var QueryData = "1h"
// var QueryMess = "\"CPU\""
// var InfluxBucket = "\"cmdb_bucket\""

// func InfluxQuery() {
// 	queryAPI := InfluxClient.QueryAPI(InfluxOrg)
// 	// get QueryTableResult
// 	sql := `from(bucket:` + InfluxBucket + `)|> range(start: -` + QueryData + `) |> filter(fn: (r) => r._measurement == ` + QueryMess + `)`
// 	sql := `from(bucket:"cmdb_bucket")|> range(start: -` + QueryData + `) |> filter(fn: (r) => r._measurement == ` + QueryMess + `)`
// 	fmt.Println(sql)
// 	// `from(bucket:"cmdb_bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "CPU")`
// 	result, err := queryAPI.Query(context.Background(), sql)
// 	if err == nil {
// 		// Iterate over query response
// 		for result.Next() {
// 			// Notice when group key has changed
// 			// if result.TableChanged() {
// 			// 	fmt.Printf("table: %s\n", result.TableMetadata().String())
// 			// }
// 			// Access data
// 			fmt.Printf("value: %v\n", result.Record().Value())
// 		}
// 		// check for an error
// 		if result.Err() != nil {
// 			fmt.Printf("query parsing error: %s\n", result.Err().Error())
// 		}
// 	} else {
// 		panic(err)
// 	}
// }

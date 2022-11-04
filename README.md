# 这是windows 客户端,用于数据采集
# 打包 windows exe 方式:go build -ldflags "-s -w -H=windowsgui"
# 查看是否启动成功,浏览器访问: localhost:9996
# windowx_export.exe 是已打包完成的执行文件

// go get github.com/shirou/gopsutil/v3/cpu@v3.22.9  获取cpu
// go get -u github.com/shirou/gopsutil/v3/load 获取负载

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

//查询数据
var QueryData = "1h"
var QueryMess = "CPU"

func InfluxQuery() {
	queryAPI := InfluxClient.QueryAPI(InfluxOrg)
	// get QueryTableResult
	// sql := "from(bucket:" + InfluxBucket + ")|> range(start: -" + QueryData + ") |> filter(fn: (r) => r._measurement == " + QueryMess + ")"
	var QueryData = "1h"
	var QueryMess = "CPU"
	sql := `from(bucket:"cmdb_bucket")|> range(start: -` + QueryData + `) |> filter(fn: (r) => r._measurement == ` + QueryMess + `)`

	fmt.Println(sql)
	// `from(bucket:"cmdb_bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "CPU")`
	result, err := queryAPI.Query(context.Background(), sql)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			// if result.TableChanged() {
			// 	fmt.Printf("table: %s\n", result.TableMetadata().String())
			// }
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}


//获取网络信息
func getNetworkInfo() {
	info, _ := net.IOCounters(true)
	for i, _ := range info {
		fmt.Printf("外出流量:%v kb\n", info[i].BytesSent%1024)
		fmt.Printf("回源流量:%v kb\n", info[i].BytesRecv%1024)
	}
}

package main

import (
	"log"
	"net/http"
	"runtime"
	"windows_export/middleware"
	"windows_export/model"

	"github.com/gin-gonic/gin"
	// "github.com/shirou/gopsutil/v3/net"
)

//获取网络信息
// func getNetworkInfo() {
// 	info, _ := net.IOCounters(true)
// 	for i, _ := range info {
// 		fmt.Printf("外出流量:%v kb\n", info[i].BytesSent%1024)
// 		fmt.Printf("回源流量:%v kb\n", info[i].BytesRecv%1024)
// 	}
// }

func routes() {
	r := gin.Default()

	//添加测试页面
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "这是windows客户端采集工具,版本:v1.0.0.0",
		})
	})

	r.Run(":9996")
}

// func TestGetProcessIdsByLocalPort(t *testing.T) {
// 	process, _ := model.GetProcess(pid)
// 	fmt.Println("ProcessName : " + string(process.szExeFile[0:]))
// 	fmt.Println("th32ModuleID : " + strconv.Itoa(int(process.th32ModuleID)))
// 	fmt.Println("ProcessID : " + strconv.Itoa(int(process.th32ProcessID)))
// 	fmt.Println("-----------------------------------------")

// 	processList := GetProcessList()
// 	fmt.Println("processList len : ", len(processList))

// 	// 因为api返回的是gbk编码的,这里对于中文就会变成乱码了
// 	// 也不要想自己转码,用这个库 "github.com/axgle/mahonia"
// 	decoder := mahonia.NewDecoder("gbk")

// 	for _, v := range processList {
// 		result := decoder.ConvertString(string(v.szExeFile[0:])) // 转为utf8
// 		fmt.Println("Process Name:", result)
// 	}
// }

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

}

func main() {
	// influxDB 初始化
	model.InfluxInit()
	// 执行定时任务
	go middleware.Cron()
	model.InfluxClose()
	printMemStats()
	// 路由
	routes()
	// influxDB 关闭
	// model.InfluxClose()
}

// func (*Process) MemoryPercent ¶
// func (p * Process ) MemoryPercent() ( float32 , error )
// MemoryPercent 返回此进程使用的总 RAM 的百分比

// go get github.com/shirou/gopsutil/v3/cpu@v3.22.9  //获取cpu
// go get -u github.com/shirou/gopsutil/v3/load //获取系统负载
// go get -u github.com/shirou/gopsutil/v3/disk  //获取磁盘
// go get -u github.com/shirou/gopsutil/v3/mem //获取内存
// go get -u github.com/robfig/cron/v3 //编写定时任务
// go get -u "github.com/axgle/mahonia" //解决编码问题

//获取主机信息
// func getHostInfo() {
// 	hostInfo, err := host.Info()
// 	if err != nil {
// 		fmt.Println("get host info fail, error: ", err)
// 	}
// 	fmt.Printf("主机名 is: %v, 主机系统: %v \n", hostInfo.Hostname, hostInfo.Platform)
// }

// go build -ldflags "-s -w -H=windowsgui"

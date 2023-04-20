package main

import (
	"net/http"
	"time"
	"windows_export/model"

	"github.com/gin-gonic/gin"
	// "github.com/shirou/gopsutil/v3/net"
)

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

//https://github.com/siruzhong/log-collection-system
//https://github.com/Darkera524/psutil_metric

// func printMemStats() {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

// }

func main() {
	// influxDB 初始化
	model.InfluxInit()
	routes()
	//每s执行一次
	for {
		//写入cpu数据
		model.GetCpuInfo()
		//上传系统负载信息
		model.GetSysLoad()
		//上传磁盘信息
		model.GetDiskInfo()
		//上传系统进程信息
		model.GetProcess()
		//上传系统内存信息
		model.GetMemInfo()

		time.Sleep(time.Second)
	}

	// 执行定时任务
	// go middleware.Cron()
	// model.InfluxClose()
	// printMemStats()

	// 路由

	// influxDB 关闭
	// model.InfluxClose()
}

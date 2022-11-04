package model

import (
	"fmt"
	"time"
	"windows_export/utils"

	"github.com/shirou/gopsutil/v3/mem"
)

//获取内存信息
func GetMemInfo() {
	data := make(map[string]interface{}, 10)

	//获取ip
	ip := utils.IP()

	//获取内存
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err: ", err)
	}
	// 获取总内存大小，单位GB
	data["memTotal"] = memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	data["memUsed"] = memInfo.Used / 1024 / 1024
	// 可用内存大小
	data["memAva"] = memInfo.Available / 1024 / 1024
	// 内存可用率
	data["memUsedPercent"] = memInfo.UsedPercent
	//写入influxdb
	InfluxWrite("System", "memory", ip, "memorys", data)
	// fmt.Println("内存信息上传完毕")
	//清除map
	delete(data, "memTotal")
	delete(data, "memUsed")
	delete(data, "memAva")
	delete(data, "memUsedPercent")
	//当前时间
	Date := time.Now()
	str := "内存信息上传完毕"
	fmt.Printf("%v %v\n", Date, str)
}

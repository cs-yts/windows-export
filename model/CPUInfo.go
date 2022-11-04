package model

import (
	"fmt"
	"log"
	"time"
	"windows_export/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

// 获取cpu信息
func GetCpuInfo() {
	//获取ip
	// GetIp := utils.IP()
	// ip := "\"" + GetIp + "\""
	ip := utils.IP()

	//获取cpu使用率
	cpuPercent, _ := cpu.Percent(time.Second, true)
	// middleware.InfluxWrite("System", "system_load", "172.16.7.10", "system_load", cpuPercent[0])
	InfluxWrite("System", "cpu_usage", ip, "cpu_usages", cpuPercent[0])
	// fmt.Println("cpu信息上传完毕 ==", cpuPercent[0])
	//当前时间
	Date := time.Now()
	str := "cpu信息上传完毕"
	fmt.Printf("%v %v\n", Date, str)
}

// 获取负载信息
func GetSysLoad() {
	//获取ip
	// GetIp := utils.IP()
	// ip := "\"" + GetIp + "\""
	ip := utils.IP()

	//获取系统负载
	loadInfo, err := load.Avg()
	if err != nil {
		log.Fatal("get average load fail. err: ", err)
	}
	InfluxWrite("System", "system_load", ip, "system_loads", loadInfo)
	// InfluxWrite("System", "system_load", "172.16.7.10", "system_load", loadInfo)
	// fmt.Println("负载信息上传完毕==", loadInfo)
	//当前时间
	Date := time.Now()
	str := "负载信息上传完毕"
	fmt.Printf("%v %v\n", Date, str)
}

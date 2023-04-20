package model

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

//存储进程数据
type Pro struct {
	Pid    int `json:"pid"`    //进程ID
	Memory int `json:"memory"` //内存使用率
}

func GetProcess() {
	// ip := utils.IP()
	// 存储进程ID及内存
	var Pro1 []Pro

	var data = make(map[string]interface{}, 10)

	// 获取进程id及对应的内存使用率
	p, _ := process.Processes()
	for _, pro := range p {
		var u1 Pro
		memory, _ := pro.MemoryInfo()
		// data["memory"] = memory.RSS
		// fmt.Println(memory)
		if memory != nil {
			u1.Pid = int(pro.Pid)
			res := (memory.RSS / 8 / 1024)
			u1.Memory = int(res)
		}

		Pro1 = append(Pro1, u1)
	}

	fmt.Println(Pro1)

	for _, v := range Pro1 {
		fmt.Println(v.Memory)
	}

	// InfluxWrite("System", "system_process", ip, "system_pro", data)
	//清除map
	delete(data, "Pid")
	delete(data, "Name")
	delete(data, "memory")

	//当前时间
	// 	Date := time.Now()
	// 	str := "进程信息上传完毕"
	// 	fmt.Printf("%v %v\n", Date, str)
}

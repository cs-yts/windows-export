package model

import (
	"fmt"
	"time"
	"windows_export/utils"

	"github.com/shirou/gopsutil/v3/disk"
)

//获取硬盘存储信息
func GetDiskInfo() {
	data := make(map[string]interface{}, 10)
	// monster = make([]map[string]string, 2)

	ip := utils.IP()
	diskPart, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(diskPart)
	for _, dp := range diskPart {
		diskUsed, _ := disk.Usage(dp.Mountpoint)

		//磁盘盘符
		data["DiskLetter"] = diskUsed.Path
		//磁盘总容量
		data["DiskTotal"] = diskUsed.Total / 1024 / 1024
		//磁盘使用率
		data["DiskUsed"] = diskUsed.UsedPercent
		InfluxWrite("System", "system_disk", ip, "system_disks", data)
		//上传完成后清理map,不然造成内存使用过大
		delete(data, "DiskLetter")
		delete(data, "DiskTotal")
		delete(data, "DiskUsed")
	}

	//当前时间
	Date := time.Now()
	str := "磁盘信息上传完毕"
	fmt.Printf("%v %v\n", Date, str)
}

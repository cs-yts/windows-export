package middleware

import (
	"log"
	"windows_export/model"

	"github.com/robfig/cron/v3"
)

func Cron() {
	//初始化定时任务
	c := cron.New(cron.WithSeconds())
	//添加定时任务
	//获取cpu
	// _, err := c.AddFunc("*/10 * * * * *", model.GetCpuInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//获取内存
	// _, err = c.AddFunc("*/10 * * * * *", model.GetMemInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//获取磁盘实例
	// _, err = c.AddFunc("*/10 * * * * *", model.GetDiskInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//获取进程
	_, err := c.AddFunc("*/10 * * * * *", model.GetProcess)
	if err != nil {
		log.Fatal(err)
	}
	//获取负载
	// _, err = c.AddFunc("*/10 * * * * *", model.GetSysLoad)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	c.Start()
	select {}
}

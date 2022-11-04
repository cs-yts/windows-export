package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//建立结构体接收我们配置文件参数
var (
	InfluxDBUrl   string
	InfluxdbToken string
	InfluxBucket  string
	InfluxOrg     string
)

//获取influxDB的配置
func InfluxDBInfo(cfg *ini.File) {
	//参数获取语法:cfg.Section("分区").key("key名").MustString("默认值")
	InfluxDBUrl = cfg.Section("influxDB").Key("InfluxDBUrl").MustString("http://127.0.0.1:8086")
	InfluxdbToken = cfg.Section("influxDB").Key("InfluxdbToken").MustString("111")
	InfluxBucket = cfg.Section("influxDB").Key("InfluxBucket").MustString("cmdb_bucket")
	InfluxOrg = cfg.Section("influxDB").Key("InfluxOrg").MustString("cmdb")
}

//初始化配置参数
func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误错误,err=", err)
	}
	//获取参数
	InfluxDBInfo(cfg)
}

# 这是windows 客户端,用于数据采集
- 打包 windows exe 方式:go build -ldflags "-s -w -H=windowsgui"
- 查看是否启动成功,浏览器访问: localhost:9996
- windowx_export.exe 是已打包完成的执行文件

<!--
// go get github.com/shirou/gopsutil/v3/cpu@v3.22.9  获取cpu
// go get -u github.com/shirou/gopsutil/v3/load 获取负载


//获取网络信息
func getNetworkInfo() {
	info, _ := net.IOCounters(true)
	for i, _ := range info {
		fmt.Printf("外出流量:%v kb\n", info[i].BytesSent%1024)
		fmt.Printf("回源流量:%v kb\n", info[i].BytesRecv%1024)
	}
} -->

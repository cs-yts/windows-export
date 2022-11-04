package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

//获取外网IP
func GetExternalIp() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

//获取内网IP(IPV4)
func GetLocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

//返回ip
func IP() string {
	//外网ip
	var InternetIP string
	//内网ip
	var IntranetIP string

	//获取外网ip
	InternetIP = GetExternalIp()

	//获取内网ip
	GetLoaclIP, _ := GetLocalIPv4s()
	for _, v := range GetLoaclIP {
		//判断前缀,排除169，127开头的值
		if !strings.HasPrefix(v, "169") {
			if !strings.HasPrefix(v, "127") {
				IntranetIP = v
			}
		}
	}

	//当外网ip跟内网ip同时存在的时候,返回外网ip
	if InternetIP != "" {
		return InternetIP
	} else {
		return IntranetIP
	}
}

package utils

import "net"

//	GetIps
//	@description: 获取主机的网卡IP地址
//	@return []string
func GetIps() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	var ips []string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				//fmt.Println(ip)
				ips = append(ips, ip)
			}
		}
	}
	return ips
}

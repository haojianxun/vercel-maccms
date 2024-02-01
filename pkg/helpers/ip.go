package helpers

import (
	"fmt"
	"net"
)

func IpToInt(ipStr string) (uint32, error) {
	// 解析 IPv4 地址
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("无效的 IP 地址")
	}

	// 确保是 IPv4 地址
	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("非 IPv4 地址")
	}

	// 将 IPv4 地址转换为整数
	ipInt := uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
	return ipInt, nil
}

func IntToIP(ipInt uint32) net.IP {
	// 将整数表示的 IP 地址转换为四个字节的形式
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt >> 24 & 0xFF)
	ip[1] = byte(ipInt >> 16 & 0xFF)
	ip[2] = byte(ipInt >> 8 & 0xFF)
	ip[3] = byte(ipInt & 0xFF)
	return ip
}

package wol

import (
	"encoding/hex"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Pc struct {
	Mac           string // MAC地址
	BroadcastAddr string // 广播地址
	Ip            string // 本机内网IP地址
}

// getLocalIP 获取本机内网IP地址
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no IP address found")
}

// getBroadcastAddr 根据本机IP地址获取广播地址
func getBroadcastAddr(ip string) string {
	ipParts := strings.Split(ip, ".")
	ipParts[3] = "255"
	return strings.Join(ipParts, ".")
}

// NewPc 构造函数，用于初始化Pc结构体
func NewPc(mac string) (*Pc, error) {
	ip, err := getLocalIP()
	if err != nil {
		return nil, err
	}
	broadcastAddr := getBroadcastAddr(ip)
	return &Pc{
		Mac:           mac,
		BroadcastAddr: broadcastAddr,
		Ip:            ip,
	}, nil
}

// WakeOnLan 发送WOL魔术包以唤醒计算机
func (p *Pc) WakeOnLan(macAddress string, broadcastAddr string) error {
	// 验证MAC地址格式
	macRegex := regexp.MustCompile(`^([A-Fa-f0-9]{2}([-:]){0,1}){5}[A-Fa-f0-9]{2}$`)
	if !macRegex.MatchString(macAddress) {
		return fmt.Errorf("invalid MAC address format")
	}

	// 移除MAC地址中的分隔符
	macAddress = strings.ReplaceAll(macAddress, ":", "")
	macAddress = strings.ReplaceAll(macAddress, "-", "")

	// 创建魔术包
	macBytes, err := hex.DecodeString(strings.Repeat(macAddress, 16))
	if err != nil {
		return err
	}
	packet := append([]byte{255, 255, 255, 255, 255, 255}, macBytes...)

	// 发送魔术包
	conn, err := net.Dial("udp", broadcastAddr+":9")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

// // 关机
func Shutdown() error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("shutdown", "/s", "/t", "0")
	case "linux", "darwin":
		cmd = exec.Command("shutdown", "-h", "now")
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
	return cmd.Run()
}

//
//

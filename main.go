package main

import (
	"HomePC-wol-mi/wol"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	client := wol.ConnectBemfa()
	defer client.Disconnect(250)

	// 设置一个信号通道来捕获中断信号，以便于程序可以优雅地退出
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to exit")
	<-sigc
	fmt.Println("Exiting...")
}

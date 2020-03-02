package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败...")
		return
	}
	fmt.Println("连接成功...")
	// 功能一：客户端可以发送单行数据，然后就输出
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入，即终端输入
	for {
		// 从终端取出用户输入的一行数据
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader.ReadString err", err)
		}
		line = strings.Trim(line, "\r\n")
		// exit 退出..
		if line == "exit" {
			fmt.Println("成功退出")
			return
		}
		// 在将line发送给服务器
		n, err := conn.Write([]byte(line + "\r\n"))
		if err != nil {
			fmt.Println("conn.Write 发送失败", err)
		}
		fmt.Printf("客户端发送%d个字符,并退出", n)
	}
}

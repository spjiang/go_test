package main

import (
	"fmt"
	"net"
)

func Process(c net.Conn) {
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1、等待客户端通过conn发送信息
		// 2、如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Println("服务器等待客户端输入:" + c.RemoteAddr().String())
		n, err := c.Read(buf)
		if err == nil {
			fmt.Println("Process Read err", err) // 从conn读取数据
			return
		}
		// 3.显示客户端发送的内容到服务器,buf[:n] 是获取buf实际的数据长度
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("s服务器开始监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务器监听失败")
	}
	fmt.Println(listen)
	// 延时关闭
	defer listen.Close()
	for {
		// 等待客户端连接
		fmt.Println("服务器等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器Accept err", err)
		} else {
			fmt.Println("服务器conn", conn)
		}
		// 这里准备用协程进行对客户端处理
		go Process(conn)
	}
}

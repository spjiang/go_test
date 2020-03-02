package main

import (
	"fmt"
	"net"
	"test/chatroom/server/model"
	"time"
)

// 这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	// 这里的pool本身就是一个全局的变量
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	// 当服务启动时，我们就去初始化我们的redis连接池，初始化pool全局变量
	initPool("localhost:6379", 16, 0, 300*time.Second)
	// 初始化UserDao实例，方便于后续直接可以使用UserDao实例
	initUserDao()
}

func main() {
	// 提示信息
	fmt.Println("服务器在监听8889端口...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	// 延时关闭监听
	defer listen.Close()

	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	// 一旦监听成功，就等待客户端连接服务器
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦连接成功，就启动一个协程和客户端保存通讯
		go process(conn)
	}
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 读出客户端发送的信息
	defer conn.Close()
	// 这里调用一个总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("客服端和服务端通讯协程错误，err=", err)
		return
	}
}

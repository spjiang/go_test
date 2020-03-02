package main

import (
	"fmt"
	"os"
	"test/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	// 接受用户选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("--------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择（1-3）：")
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室...")
			fmt.Println("请输入用户Id:")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("请输入用户密码:")
			_, _ = fmt.Scanln(&userPwd)
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("Login result err:", err)
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户Id:")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("请输入用户密码:")
			_, _ = fmt.Scanln(&userPwd)
			fmt.Println("请输入用户名称:")
			_, _ = fmt.Scanln(&userName)
			up := &process.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("Login result err:", err)
			}
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}

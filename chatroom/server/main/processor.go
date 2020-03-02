package main

import (
	"fmt"
	"io"
	"net"
	message "test/chatroom/common"
	process2 "test/chatroom/server/process"
	"test/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个serverProcessMes函数
// 功能：根据客服端发送的信息的总类不通，决定调用那个函数来处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err := up.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println("serverProcessLogin fail err", err)
		}
	case message.RegisterMesType:
		// 处理注册
		up := &process2.UserProcess{
			Conn: p.Conn,
		}
		err := up.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println("ServerProcessRegister fail err", err)
		}
	case message.SmsMesType: // 群聊
		up := &process2.SmsProcess{}
		err := up.SendGroupMes(mes)
		if err != nil {
			fmt.Println("SendGroupMes fail err", err)
		}
	default:
		fmt.Println("消息类型不存在，暂时无法处理...")
	}
	return
}

func (p *Processor) Process2() (err error) {
	for {
		fmt.Println("读取客户端发送的信息...")
		//这里我们将读取的数据包。直接封装成一个函数readPkg()
		// conn.read 在conn没有被关闭的情况下，才会阻塞
		// 如果客服端关闭了，conn，则，就不会阻塞
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出...")
				return err
			} else {
				fmt.Println("pkg err=", err)
				return err
			}
		}
		fmt.Println("mes=", mes)
		// 处理信息
		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes fail err=", err)
			return err
		}
	}
}

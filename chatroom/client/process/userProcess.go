package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"test/chatroom/client/utils"
	message "test/chatroom/common"
)

type UserProcess struct {
}

func (u *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1、连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		return errors.New("net.Dial fail")
	}

	// 延时关闭
	defer conn.Close()

	// 2、准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3、创建一个LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 4、将loginMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		return err
	}

	// 5、把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6、将mes进行序列化
	mesBody, err := json.Marshal(mes)
	if err != nil {
		return err
	}

	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(mesBody)
	if err != nil {
		fmt.Println("注册信息发送失败...err=", err)
	}

	responseMes, err := tf.ReadPkg()
	if err != nil {
		return err
	}
	// 将mes的data数据反系列化成LoginResMes
	var registerResMes message.RegisterResMes

	err = json.Unmarshal([]byte(responseMes.Data), &registerResMes)
	if err != nil {
		return err
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return nil
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	// 下面我们要定一个协议
	// 1、连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		return errors.New("net.Dial fail")
	}

	// 延时关闭
	defer conn.Close()

	// 2、准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3、创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4、将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		return err
	}

	// 5、把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6、将mes进行序列化
	mesBody, err := json.Marshal(mes)
	if err != nil {
		return err
	}

	// 7、到这个时候，data就是我们要发送的信息
	//	7.1、先把data的长度发送给服务器
	//   先获取到data长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(mesBody))

	// 定义一个byte数组用于把pkgLen数字转化切片进行存储
	var buf [4]byte
	// 数组是引用类型，更改了切片就直接更改了bytes数组
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		return err
	}

	// 发送内容
	_, err = conn.Write(mesBody)
	if err != nil {
		return err
	}

	// fmt.Printf("\n客户端，发送信息的长度=%d，内容%s\n", len(mesBody), string(mesBody))

	// 这里还需要处理服务器返回的信息.
	// 读取数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	responseMes, err := tf.ReadPkg()
	if err != nil {
		return err
	}
	// 将mes的data数据反系列化成LoginResMes
	var loginResMes message.LoginResMes

	err = json.Unmarshal([]byte(responseMes.Data), &loginResMes)
	if err != nil {
		return err
	}
	if loginResMes.Code == 200 {

		//初始化
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline


		// 这里我们还需要在客户端启动一个协程
		// 该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		// 则接受并显示在客户端的终端
		// 可以显示当前在线用户列表，遍历
		fmt.Println("当前在线用户列表如下：")

		for _, v := range loginResMes.UserIds {
			// 如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}
			fmt.Println("当前用户id：\t", v)
			// 完成客户端的onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		// 处理和服务端进行的通信
		go serverProcessMes(conn)

		// 1、显示我们的登录成功菜单.
		for {
			ShowMenu()
		}
	} else {
		return errors.New(loginResMes.Error)
	}
	return
}

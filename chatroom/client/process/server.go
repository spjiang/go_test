package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"test/chatroom/client/utils"
	message "test/chatroom/common"
)

// 显示成功登录的界面
func ShowMenu() {
	fmt.Println("----------------恭喜XXX登录成功------------------")
	fmt.Println("\t\t 1.显示在线用户列表")
	fmt.Println("\t\t 2.发送消息")
	fmt.Println("\t\t 3.信息列表")
	fmt.Println("\t\t 4.退出系统")
	fmt.Println("请选择（1-4）：")
	var key int

	// 发送消息内容
	var content string

	_, err := fmt.Scanln(&key)
	if err != nil {
		fmt.Println("读取指令错误")
	}

	// 因为，我们总会使用到SmsProcess实例，因此我们将其定义到switch外部

	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说点什么呢？")
		_, err := fmt.Scanln(&content)
		if err != nil {
			fmt.Println("读取信息失败")
		}
		err = smsProcess.SendGroupMes(content)
		if err != nil {
			fmt.Println("发送消息失败")
		}

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统...")
		os.Exit(0)
	}
}

// 和服务器保存通讯
func serverProcessMes(conn net.Conn) {
	// 创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端正在等待读取服务器发送的消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType: // 用户通知, 有人上线了
			//1、取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes):", err)
			}
			//2、把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: // 有人群发消息了
			outputGroupMes(&mes)
		}
		// 如果读取信息，就直接展示出来
		fmt.Printf("mes=%v\n", mes)
	}
}

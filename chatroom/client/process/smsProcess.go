package process

import (
	"encoding/json"
	"fmt"
	"test/chatroom/client/utils"
	message "test/chatroom/common"
)

type SmsProcess struct {
}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	// 2创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err=", err)
		return
	}

	mes.Data = string(data)
	//4再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err=", err)
		return
	}
	//5将mes发送给服务器
	tf := utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes tf.WritePkg(data) err=", err)
		return
	}
	return
}

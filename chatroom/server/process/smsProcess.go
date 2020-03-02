package process2

import (
	"encoding/json"
	"fmt"
	"net"
	message "test/chatroom/common"
	"test/chatroom/server/utils"
)

type SmsProcess struct {
}

// 写一个方法进行转发信息
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 遍历服务器的onlineUsers map[int]*UserProcesss
	// 将消息转发取出
	// 取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal([]byte(mes.Data), &smsMes) err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes data,err := json.Marshal(mes):err=", err)
		return
	}
	for _, up := range userMgr.onlineUsers {
		if smsMes.UserId == up.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser tf.WritePkg(data) err=", err)
	}
}

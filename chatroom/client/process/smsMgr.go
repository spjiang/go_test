package process

import (
	"encoding/json"
	"fmt"
	message "test/chatroom/common"
)

func outputGroupMes(mes *message.Message) { // 这个地方一定是smsMes
	// 显示即可
	// 1反序化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal([]byte(mes.Data),&smsMes) err=", err)
		return
	}
	// 显示信息
	info := fmt.Sprintf("用户ID：%d\t对大家说：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}

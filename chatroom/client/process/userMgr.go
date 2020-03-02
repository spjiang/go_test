package process

import (
	"fmt"
	"test/chatroom/client/model"
	message "test/chatroom/common"
)

var onlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser


// 编写一个方法，处理返回的 NotifyUserStatus
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	// 遍历一把onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户ID=：", id)
	}
}

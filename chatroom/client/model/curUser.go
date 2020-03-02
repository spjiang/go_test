package model

import (
	"net"
	message "test/chatroom/common"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

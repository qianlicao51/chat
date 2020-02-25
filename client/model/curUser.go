package model

import (
	"net"
	"chat/common/message"
)

// 当前用户|初始化位置>
type CurUser struct {
	Conn net.Conn
	message.User
}

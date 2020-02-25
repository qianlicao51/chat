package model

import (
	"chat/common/message"
	"net"
)

// 当前用户|初始化位置>
type CurUser struct {
	Conn net.Conn
	message.User
}

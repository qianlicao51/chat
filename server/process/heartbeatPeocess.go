package processSer

import "fmt"

//服务器心跳检测客户端是否还在
type HeartBeat struct {
}

func (h *HeartBeat) CheckClientIsOnline() {
	for i := range userMgr.onlineUsers {
		fmt.Println("在线用户id", i)
	}
	fmt.Println()
}

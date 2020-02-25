package processSer

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"github.com/huandu/xstrings"
)

//服务器心跳检测客户端是否还在
type HeartBeatProcess struct {
}

//心条检测|服务器有客户端的conn,如果数据发送不出去就是 客户端失去联系，不需要等待客户端回应
func (h *HeartBeatProcess) HeartBeatRequest() {
	for userid, process := range userMgr.onlineUsers {
		fmt.Println(userid, process)
		tr := &utils.Transfer{
			Conn: process.Conn,
		}
		var mes message.Message
		mes.Type = message.HeartBeatMesType
		mes.Data = ""
		mesData, err := json.Marshal(mes)
		if err != nil {
			fmt.Println("心跳检测信息序列化错误")
			break
		}
		if err = tr.WritePkg(mesData); err != nil {
			//发送不出去心跳表示用户下线了
			fmt.Println("用户ID 心跳检测err:", userid)
			userMgr.DelOnlineUser(userid)
		}
	}
	fmt.Println(xstrings.Center("检测之后结果", 30, "*"))
	userMgr.ShowOnlineUsser()
	fmt.Println()
}

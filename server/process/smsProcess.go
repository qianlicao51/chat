package processSer

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//[暂时不需要字段]
}

// 转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	//遍历服务器的onlineUsers
	//将消息转发出去
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("发送消息~反序列化失败~", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes 序列化失败,", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if smsMes.UserID == id {
			continue
		}
		this.SendMesEachOnlineUser(data, up.Conn)
	}
	return
}

// 发送消息给除了自己的其他人
func (this *SmsProcess) SendMesEachOnlineUser(info []byte, conn net.Conn) {
	tr := &utils.Transfer{
		Conn: conn,
	}
	tr.WritePkg(info)
}

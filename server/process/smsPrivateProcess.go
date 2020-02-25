package process2

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//私聊
type SmsPrivateProcess struct {
	//[暂时不需要字段]
}

// 转发消息
func (this *SmsPrivateProcess) SendPirvateMes(mes *message.Message) (err error) {

	//将私聊消息转发出去
	var smsMes message.SmsPrivateMes //此变量存在的意义:从中获取 私聊 对方ID
	err = json.Unmarshal([]byte(mes.Data), &smsMes)

	privateUserUp, ok := userMgr.onlineUsers[smsMes.ChatUserID]
	if !ok {
		fmt.Println("私聊 查找不到 用户(id)", smsMes.ChatUserID)
		return errors.New("查找不到 私聊用户")
	}
	if err != nil {
		fmt.Println("发送消息~反序列化失败~", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes 序列化失败,", err)
		return
	}
	this.SendMesPrivate(data, privateUserUp.Conn)
	return
}

// 发送消息(私聊)
func (this *SmsPrivateProcess) SendMesPrivate(info []byte, conn net.Conn) {
	tr := &utils.Transfer{
		Conn: conn,
	}
	tr.WritePkg(info)
}

package process

import (
	"chat/client/utils"
	"chat/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	// 2 创建SmsMes
	var smsMes message.SmsMes
	smsMes.Context = content //内容

	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus
	//3 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}
	mes.Data = string(data)
	//4 对mes序列化

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}
	//5 将mes发送服务器
	tr := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("sendMes ERR=", err)
		return
	}
	return
}

//发送私聊消息
func (this *SmsProcess) SendPrivateMes(userId int, content string) (err error) {
	//判断是否是在线用户
	_, ok := onlineUser[userId]
	if !ok {
		fmt.Println("当前用户不在线|没有此用户")
		return nil
	}
	//1 创建一个Mes|私聊
	var mes message.Message
	mes.Type = message.SmsPrivateMesType

	// 2 创建SmsMes
	var smsMes message.SmsPrivateMes
	smsMes.Context = content //内容

	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus
	smsMes.ChatUserID = userId //私聊对方ID
	//3 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}
	mes.Data = string(data)
	//4 对mes序列化

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}
	//5 将mes发送服务器
	tr := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("sendMes ERR=", err)
		return
	}
	return
}

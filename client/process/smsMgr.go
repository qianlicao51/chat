package process

import (
	"chat/common/message"
	"encoding/json"
	"fmt"
)

/********************************************
			处理群发消息
*********************************************/

//显示消息
func outPutGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("客户端得到 消息 反序列化失败,", err)
		return
	}
	//显示信息
	showContext := fmt.Sprintf(">> 用户ID%d 对大家说:\t%s", smsMes.UserID, smsMes.Context)
	fmt.Println(showContext)
	fmt.Println()
}

//显示消息|私聊消息
func outPutPrivateMes(mes *message.Message) {
	var smsMes message.SmsPrivateMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("客户端得到 消息 反序列化失败,", err)
		return
	}
	//显示信息
	showContext := fmt.Sprintf(">> 用户ID%d 对你说:\t%s", smsMes.UserID, smsMes.Context)
	fmt.Println(showContext)
	fmt.Println()
}

package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"chat/client/utils"
	"chat/common/message"
	"strconv"
)

//保持和服务器端的通信
//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("-----登录成功------")
	fmt.Println("\t\t\t1 显示在线用户列表")
	fmt.Println("\t\t\t2 发送消息(群发)")
	fmt.Println("\t\t\t3 信息列表")
	fmt.Println("\t\t\t4 退出系统")
	fmt.Println("\t\t\t5 发送文件到服务器")
	fmt.Println("\t\t\t6 发送消息(私聊)")
	fmt.Println("---------选择1-4---------")

	var key int
	var content string
	//因为总会使用到SmsProcess，创建在此处复用
	smsProcess := &SmsProcess{}
	osIn := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("输入给大家发送的消息！！")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("即将退出系统")
		os.Exit(0)

	case 5:
		fmt.Println("发送文件到服务器")
		fmt.Scanf("%s\n", &content)
		sendFileServer := &SendFileToServer{}
		sendFileServer.sendFile(content)

	case 6:
		//显示在线用户
		outputOnlineUser()
		fmt.Println("输入要私聊的用户ID")
		readLine, _, _ := osIn.ReadLine()
		sendToUserId, err := strconv.Atoi(string(readLine))
		if err!=nil {
			fmt.Println("输入用户的id错误")
			return
		}
		fmt.Println("输入要私聊的内容")
		sendContent, _, _ := osIn.ReadLine()
		smsProcess.SendPrivateMes(sendToUserId,string(sendContent))
	default:
		fmt.Println("输入有误重新输入！")
		fmt.Println()
	}
}

//和服务器端保持通信
func serverProcessMes(conn net.Conn) {
	//不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8064]byte{},
	}
	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("服务器端出现错误 ", err)
			return
		}
		//	读取到消息，进行下一步处理
		fmt.Printf("msg=%v\n", mes)
		//	读取消息，逻辑处理
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//	有人上线
			/**
			1 取出NotifyUserStatusMes
			2 把这个用户保存到客户端维护的map[int]User中
			*/
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("在线列表回显失败", err)
			}
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsMesType: //群发消息
			outPutGroupMes(&mes)
		case message.SmsPrivateMesType: // 私聊消息
			outPutPrivateMes(&mes)
		default:
			fmt.Println("服务器返回的是未知类型")
		}
	}
}

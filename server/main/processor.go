package main

import (
	"chat/common/message"
	process2 "chat/server/process"
	"chat/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据发送消息种类不同，处理不同调用
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	//客户端发来的数据
	//fmt.Println("--客户端发来的数据:", mes.Type)
	switch mes.Type {
	case message.LoginMesType:
		fmt.Println("客户端登录")
		up := &process2.UserProcess{Conn: this.Conn}
		up.ServerProcessLogin(mes)
		//err = process.ServerProcessLogin(conn, mes)
	case message.LoginResMesType:
	//登录处理
	case message.RegisterMesType:
		//注册
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//转发消息，完成群聊
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SendFileMesType:
		//文件传输服务
		fmt.Println("接受到客户端发送来的文件")
		fileProcess := &process2.SendFileProcess{Conn: this.Conn}
		fileProcess.ReceiveFile(mes)

	case message.SmsPrivateMesType:
		//私聊服务
		smsPrivate := &process2.SmsPrivateProcess{}
		fmt.Println("服务器接受到的私聊信息~", mes)
		smsPrivate.SendPirvateMes(mes)
	default:
		fmt.Println("未知类型，无法处理")
		fmt.Println()
	}
	return
}
func (this *Processor) Process2() (err error) {
	//循环的客户端发送的消息
	for {
		//创建transform
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了，服务器也正常退出")
				return err
			} else {
				fmt.Println("未知错误", err)
				return err
			}
		}
		//fmt.Println(mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println(" server 退出")
			return err
		}
	}
	return nil
}

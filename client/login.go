package main

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//登陆校验
func Login(userId int, userPwd string) (err error) {

	//fmt.Println(userId, userPwd)
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接服务器失败", err)
		return
	}
	//通过conn消息发送给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//消息结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// loginMes序列化
	//mes.Data=json.Marshal()

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	//data赋值给 mes.Data
	mes.Data = string(data)

	//mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	//此时 data就是发送消息|先发送 data的长度到服务器|
	//获取data长度，转为表示长度的切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen) //长度转为 byte

	defer conn.Close() //记得延时关闭
	n, err := conn.Write(bytes[0:4])
	if err != nil || n != 4 {
		fmt.Println("发送失败", err)
		return
	}
	//fmt.Println("客户端发送消息 长度成功！",len(data),"发送的数据",string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write data fial.", err)
		return
	}
	//接受服务器返回的信息

	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8064]byte{},
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取失败1", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功~")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

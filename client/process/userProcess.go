package process

import (
	"chat/client/utils"
	"chat/common/iniconst"
	"chat/common/message"
	utils2 "chat/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

//关联用户登录的方法

//登陆校验
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//连接到服务器
	serverAddr := utils2.ConfGetValString(iniconst.INI_NAME_CLIENT, iniconst.CLIENT_TO_SER_ADDR)
	fmt.Println("服务器地址：", serverAddr)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("客户端连接服务器失败", err)
		return
	}
	defer conn.Close() //记得延时关闭
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

	n, err := conn.Write(bytes[0:4])
	if err != nil || n != 4 {
		fmt.Println("发送(数据长度)失败", err)
		return
	}
	fmt.Println("客户端发送消息 长度成功！", len(data), "发送的数据", string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("client  conn.write data fial.", err)
		return
	}
	//接受服务器返回的信息
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8064]byte{},
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取服务器数据失败", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	//登录成功
	if loginResMes.Code == 200 {
		//初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserID = loginMes.UserId
		CurUser.UserName = loginMes.UserName
		CurUser.UserStatus = message.UserOnLice

		//显示在线用户
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UserIds {
			if v == userId { //不显示当前自己
				continue
			}
			fmt.Printf("用户id :%v\n", v)
			//初始化客户端维护的 在线列表
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnLice,
			}
			onlineUser[userId] = user

		}
		fmt.Println("-----------------------")

		//	显示登陆成功的菜单
		//开启协程，该协程和服务器有数据交互，如果服务器有数据推送给客户端，则接受显示
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Regist(userId int, userPwd, userName string) (err error) {
	serverAddr := utils2.ConfGetValString(iniconst.INI_NAME_CLIENT, iniconst.CLIENT_TO_SER_ADDR)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("客户端连接服务器失败", err)
		return
	}
	defer conn.Close() //记得延时关闭
	//通过conn消息发送给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//消息结构体
	var registerMes message.RegisterMes
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserID = userId
	registerMes.User.UserName = userName
	// loginMes序列化
	//mes.Data=json.Marshal()

	data, err := json.Marshal(registerMes)
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

	//发送 数据 到服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错！", err)
	}
	//读取返回信息

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("读取失败", err)
		return
	}
	//
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		//注册成功，重新登录
		fmt.Println("注册成功，重新登录")
	} else {
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return
}

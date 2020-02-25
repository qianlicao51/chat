package main

import (
	"chat/common/iniconst"
	"chat/server/model"
	processSer "chat/server/process"
	"chat/utils"
	"fmt"
	"github.com/huandu/xstrings"
	"log"
	"net"
	"time"
)

func main() {
	//初始化Redis连接池
	redisAddr := utils.ConfGetValString(iniconst.INI_NAME_REDIS, iniconst.REDIS_ADDR)
	initPool(redisAddr)
	//初始化userdao
	initUserDao()
	//	服务器 监听8889端口
	serverAddr := utils.ConfGetValString(iniconst.INI_NAME_SERVER, iniconst.SERVER_ADDR)
	log.Println(xstrings.Center("服务器地址:"+serverAddr, 30, "-"))
	listen, err := net.Listen("tcp", serverAddr)
	defer listen.Close()
	if err != nil {
		fmt.Println("server listen error！")
		return
	}

	go heartBeatCheck()
	//端口监听成功，等待客户端连接服务器
	for {
		log.Println("等待客户端连接……")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("server listen.accept err=", err)
		}
		//连接成功，启动协程和客户端保持通信
		go process(accept)
	}
}

//服务端心跳检测
func heartBeatCheck() {
	beatProcess := &processSer.HeartBeatProcess{}
	for {
		time.Sleep(20 * time.Second)
		log.Println("服务器心跳检测判断客户端是否仍然在线")
		beatProcess.HeartBeatRequest()
	}
}

//初始化userdao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)

}

//处理和客户端的通信
func process(conn net.Conn) {
	defer conn.Close()
	pr := &Processor{Conn: conn}
	err := pr.Process()
	if err != nil {
		log.Println("go 协程出现错误，退出", err)
		//这里应该是用户CTRL+C 退出了
		return
	}
}

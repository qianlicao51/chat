package processSer

import (
	"chat/common/message"
	"chat/server/model"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
	//表示该conn哪个用户的
	UserId int
}

//用户相关 处理

//通知所有在线用户的方法|userid通知其他在线用户。
func (this *UserProcess) NotifyOthersOnliceUser(userId, userStatus int) {
	//遍历userMgr.onlineUsers ,然后一个个 发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		if id == userId { //自己不通知自己
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId, userStatus)
	}

}

//通知别人，我上线了
func (this *UserProcess) NotifyMeOnline(userId, userStatus int) {
	//	组装notify消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes

	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = userStatus //

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化用户状态出错，", err)
		return
	}
	mes.Data = string(data)
	//	mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化用户状态出错，", err)
		return
	}
	//	准备发送
	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("通知消息的时候发现错误", err)
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registMes message.RegisterMes
	//取出数据，序列化为 RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterMesType
	//返回的信息
	var registerResMes message.RegisterResMes

	user := registMes.User

	err = model.MyUserDao.Register(&user)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505 //用户已经存在
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//发送数据

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	//发送|使用了分层，先创建trans实例
	tf := &utils.Transfer{
		Conn: this.Conn,
		//Buf:  [8064]byte{},
	}
	tf.WritePkg(data)
	//_ = utils.WritePkg(conn, data)
	return

}

// 登录处理
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//取出data，并直接反序列化成loginMes
	var loginMes message.LoginMes
	fmt.Println("登录处理ServerProcessLogin")

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//返回的信息
	var loginResMes message.LoginResMes

	//TODO 判断是否已经登录了
	_, err = userMgr.GetOnlineUserByID(loginMes.UserId)
	if err != nil {
		//用户已经在线了
		loginResMes.Code = 100
		loginResMes.Error = "用户已经登录"
	} else {
		//redis数据库完成验证
		//1 获取连接
		user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
		fmt.Println("登录的用户是：", user)
		if err != nil {
			if err == model.ERROR_USER_NOTEXISTS {
				loginResMes.Code = 500
				loginResMes.Error = err.Error()
			} else if err == model.ERROR_USER_PWD {
				//密码错误
				loginResMes.Code = 300
				loginResMes.Error = err.Error()
			} else {
				loginResMes.Code = 505
				loginResMes.Error = "服务器内部错误"
			}
		} else {
			loginResMes.Code = 200
			//	将登录成功的用户放入
			this.UserId = loginMes.UserId
			userMgr.AddOnlineUser(this)
			//通知其他人我上线
			this.NotifyOthersOnliceUser(loginMes.UserId, message.UserOnLice)

			//将当前在线用户的id放入到 LoginResMes.UserIds
			for id, _ := range userMgr.onlineUsers {
				loginResMes.UserIds = append(loginResMes.UserIds, id)
			}
			fmt.Println("登录成功")
		}
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	//发送|使用了分层，先创建trans实例
	tf := &utils.Transfer{
		Conn: this.Conn,
		//Buf:  [8064]byte{},
	}
	tf.WritePkg(data)
	//_ = utils.WritePkg(conn, data)
	return
}

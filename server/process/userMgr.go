package processSer

import (
	"fmt"
	"net"
	"strings"
)

//服务器只有1个，很多地方都用，因此，定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	//存储在线用户
	onlineUsers map[int]*UserProcess
	//TODO 此位置，怎么修改为 键值对为指针|或者他本身就是指针？
	//与上面相反，因为用户意外下线可以通过 conn 判断，以此判断用户下线
	onlineUsersProcess map[net.Conn]int
}

//初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers:        make(map[int]*UserProcess, 1024),
		onlineUsersProcess: make(map[net.Conn]int, 1024),
	}
}

// add
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
	this.onlineUsersProcess[up.Conn] = up.UserId
}

//服务器显示在线用户
func (this *UserMgr) ShowOnlineUsser() {
	for _, process := range this.onlineUsers {
		fmt.Println("在线用户id", process.UserId)
	}
	fmt.Println(strings.Repeat("_", 20))
	for _, index := range this.onlineUsersProcess {
		fmt.Println("在线用户ID", index)
	}
}

// del
func (this *UserMgr) DelOnlineUser(userId int) {
	//注意这个先后顺序
	delete(this.onlineUsersProcess, this.onlineUsers[userId].Conn)
	delete(this.onlineUsers, userId)
}

//查找
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id对应
func (this *UserMgr) GetOnlineUserByID(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return
}

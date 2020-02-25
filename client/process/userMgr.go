package process

import (
	"fmt"
	"chat/client/model"
	"chat/common/message"

	"github.com/cheynewallace/tabby" //构建输出用户界面
	sxtrings "github.com/huandu/xstrings"
)

/**********************************
			全局变量
**********************************/
//客户端维护的在线map
var onlineUser map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //登录成功后，完成初始化

// 当前在线用户列表|有人登录时会及时更新
func outputOnlineUser() {
	fmt.Println(sxtrings.Center("当前在线用户列表", 30, "*"))
	t := tabby.New()
	t.AddHeader("用户ID", "用户详情")
	for id, user := range onlineUser {
		t.AddLine(id, user)
		//fmt.Println("用户id\t", id, " -|", user)
	}
	t.Print()
}

//处理返回的notif信息
func updateUserStatus(notify *message.NotifyUserStatusMes) {
	user, ok := onlineUser[notify.UserId]
	if !ok {
		user = &message.User{
			UserID:     notify.UserId,
			UserStatus: notify.UserStatus,
		}
	}
	//更新状态
	user.UserStatus = notify.UserStatus
	onlineUser[notify.UserId] = user
	outputOnlineUser()
}

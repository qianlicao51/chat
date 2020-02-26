package processSer

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"github.com/huandu/xstrings"
	"log"
	"net"
)

//服务器心跳检测客户端是否还在
type HeartBeatProcess struct {
}

//TODO 废弃这个心跳检测，使用下面那个OfflineDeals 因为下面那个是实时的判断
//心条检测|服务器有客户端的conn,如果数据发送不出去就是 客户端失去联系，不需要等待客户端回应
func (h *HeartBeatProcess) HeartBeatRequest() {
	offlineUserid := make([]int, 0)
	for userid, process := range userMgr.onlineUsers {
		fmt.Println(userid, process)
		tr := &utils.Transfer{
			Conn: process.Conn,
		}
		var mes message.Message
		mes.Type = message.HeartBeatMesType
		mes.Data = ""
		mesData, err := json.Marshal(mes)
		if err != nil {
			fmt.Println("心跳检测信息序列化错误")
			break
		}
		if err = tr.WritePkg(mesData); err != nil {
			//发送不出去心跳表示用户下线了
			fmt.Println("用户ID 心跳检测err:", userid)
			offlineUserid = append(offlineUserid, userid) //下线用户
			userMgr.DelOnlineUser(userid)
		}
	}
	fmt.Println(xstrings.Center("检测之后结果", 30, "*"))
	userMgr.ShowOnlineUsser()
	if len(offlineUserid) >= 1 {
		//通知其他人下线用户
		for _, v := range offlineUserid {
			fmt.Println("下线用户ID:", v)
		}
	}
	fmt.Println()
}

//因为意外断开 导致服务器读取不到而下线 这种情况处理
func (h *HeartBeatProcess) OfflineDeals(conn net.Conn) {
	userid, ok := userMgr.onlineUsersProcess[conn]
	fmt.Println("意外中断下线用户是userID:", userid, conn)
	if ok {
		userMgr.DelOnlineUser(userid)
		//下线通知
		process := &UserProcess{Conn: nil, UserId: userid}
		process.NotifyOthersOnliceUser(userid, message.UserOffLine)
		log.Println("发送下线通知，下线用户", userid)
	}
	userMgr.ShowOnlineUsser()
}

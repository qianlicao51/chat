package message

//user 用户结构体
type User struct {
	UserID     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态 定义在常量 UserOnLice
}

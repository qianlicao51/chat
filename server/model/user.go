package model

//user 用户结构体
type User struct {
	UserID   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

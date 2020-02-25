package message

//消息类型
const (
	LoginMesType            = "LoginMes"            //登录请求
	LoginResMesType         = "LoginResMes"         //登录响应
	RegisterMesType         = "RegisterMes"         //注册
	NotifyUserStatusMesType = "NotifyUserStatusMes" //用户上线通知
	RegisterResMesType      = "RegisterResMes"
	SmsMesType              = "SmsMes"        //群聊
	SendFileMesType         = "SendFileMes"   //客户端 向服务器发送文件
	SmsPrivateMesType       = "SmsPrivateMes" //私聊

	HeartBeatMesType    = "HeartBeatMes"    //	 服务器心跳检测
	HeartBeatResMesType = "HeartBeatResMes" //客户端相应服务器的心跳检测
)

// 定义用户状态的常量
const (
	UserOnLice = iota
	UserOffLine
	UserBusyLine
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

// 登陆消息
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"` //用户名
}

//登录响应
type LoginResMes struct {
	Code    int    `json:"code"` //状态码 500 为注册，200 登陆成功
	UserIds []int  //增加字段，保存用户ID切片
	Error   string `json:"error"` //返回错误信息
}
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}
type RegisterResMes struct {
	Code  int    `json:"code"`  //状态码 400 已经占用|，200 成功
	Error string `json:"error"` //返回错误信息
}

// 为了配合服务器端的推送通知，用户状态变化
type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`     //用户ID
	UserStatus int `json:"userStatus"` //用户状态
}

//群聊消息
type SmsMes struct {
	Context string `json:"context"`
	User           //匿名结构体
}

//私聊消息
type SmsPrivateMes struct {
	SmsMes
	ChatUserID int `json:"chatUserID"` //私聊对方ID

}

//发送文件
type SendFileMes struct {
	FileName    string //文件名称
	FileContext string //|文件内容[]byte使用base64转换
}

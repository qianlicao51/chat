package processSer

import (
	"chat/common/iniconst"
	"chat/common/message"
	utils2 "chat/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type SendFileProcess struct {
	Conn net.Conn
}

//服务器接收文件|文件使用base64编码
func (this *SendFileProcess) ReceiveFile(mes *message.Message) (err error) {
	var fileMes message.SendFileMes
	err = json.Unmarshal([]byte(mes.Data), &fileMes)
	if err != nil {
		fmt.Println("~ReceiveFile~", err)
		return
	}
	fmt.Println("传输的文件是：", fileMes.FileName)
	saveFileBaseDIR := utils2.ConfGetValString(iniconst.INI_NAME_SERVER, iniconst.SERVER_SAVEFILE)
	openFile, err := os.OpenFile(saveFileBaseDIR+utils2.GetDateStr()+fileMes.FileName, os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	defer openFile.Close()
	//发送使用了base64 编码，现在解码
	decodeString, err := base64.StdEncoding.DecodeString(fileMes.FileContext)
	if err != nil {
		fmt.Println("服务器接收文件base64解码错误", err)
		return
	}
	openFile.Write(decodeString)
	return
}

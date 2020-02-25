package process

import (
	"chat/client/utils"
	"chat/common/message"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/********************************************
			|客户端发送数据到服务器
*********************************************/
// 发送文件到服务器
type SendFileToServer struct {
}

func (this *SendFileToServer) sendFile(filePath string) (err error) {

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("输入文件路径错误", err)
		return
	}
	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	fmt.Println("发送的文件名是：", fileName)

	var mes message.Message
	mes.Type = message.SendFileMesType

	open, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer open.Close()
	var count int64
	//传输数据
	for {
		buf := make([]byte, 1024)
		//	读取文件内容
		readN, err := open.Read(buf)

		if err != nil && err == io.EOF {
			fmt.Println("文件传输完成")
			//通知服务器文件传输结束
			//conn.Write([]byte("finish"))
			break
		}
		//发送到服务端

		var fileMes message.SendFileMes
		fileMes.FileName = fileName

		//多读取到的文件进行编码
		encodeToString := base64.StdEncoding.EncodeToString(buf[:readN])
		fileMes.FileContext = encodeToString

		data, err := json.Marshal(fileMes)
		if err != nil {
			fmt.Println("文件传输序列化错误", err)
			return err
		}
		mes.Data = string(data)
		data, err = json.Marshal(mes)
		if err != nil {
			fmt.Println("文件传输序列化错误", err)
			return err
		}
		tr := &utils.Transfer{
			Conn: CurUser.Conn,
		}

		err = tr.WritePkg(data)
		count += int64(readN)

		sendPercent := float64(count) / float64(fileSize) * 100
		value := fmt.Sprintf("%.2f", sendPercent)
		//打印上传进度
		fmt.Println("文件上传：" + value + "%")
	}
	return

}

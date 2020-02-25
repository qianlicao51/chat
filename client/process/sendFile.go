package process

import (
	"chat/client/utils"
	"chat/common/message"
	utils2 "chat/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

/********************************************
			|客户端发送数据到服务器
*********************************************/
// 发送文件到服务器
type SendFileToServer struct {
}

func (this *SendFileToServer) sendFile(filePath string) (err error) {
	if strings.HasPrefix(filePath, "\"") {
		filePath = string(filePath[len("\""):]) //去掉前面的空格。参照strings.HasPrefix的写法
	}
	if strings.HasSuffix(filePath, "\"") {
		filePath = string(filePath[:len(filePath)-1]) //去掉前面的空格。参照strings.HasPrefix的写法
	}
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		fmt.Println("输入文件路径错误", err)
		return
	}
	//TODO 不显示文件进度了
	fileSize := fileInfo.Size()
	if fileSize > (50 << 20) {
		fmt.Println("文件大小大于50M 不允许发送，文件大小:", fileSize)
		return
	}
	var mes message.Message
	mes.Type = message.SendFileMesType

	open, err := os.Open(filePath)
	//TODO 错误原因:这个函数把文件读取完毕了，下面读取到的是EOF，使用Seek重新移动 句柄(句柄这个词或许不准确)
	fileSha1 := utils2.FileSha1(open)
	open.Seek(0, 0)

	fileName := fileSha1 + "_" + fileInfo.Name()
	fmt.Println("开始发送文件:", fileName)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer open.Close()
	var count int64
	//传输数据
	startData := utils2.GetDateStrHmS()
	for {
		buf := make([]byte, 1024*4)
		readN, err := open.Read(buf) // 	读取文件内容
		if err != nil && err == io.EOF {
			fmt.Println("文件传输完成", err)
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
		if err != nil {
			fmt.Println("发送文件|服务器端连接失败", err)
			break
		}
		count += int64(readN)

		//TODO 不显示文件进度了
		//sendPercent := float64(count) / float64(fileSize) * 100
		//value := fmt.Sprintf("%.2f", sendPercent)
		////打印上传进度
		//fmt.Println("文件上传：" + value + "%")
	}
	endData := utils2.GetDateStrHmS()
	fmt.Println("文件发送 开始和完毕时间 分别如下")
	fmt.Println(startData)
	fmt.Println(endData)
	return

}

package utils

import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8064]byte //传输时使用的缓冲
}

//http://lihaoquan.me/2016/11/5/golang-byteorder.html|go语言的字节序
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//读取数据包封装为函数|readPkg()
	fmt.Println("读取客户端发来的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("读取失败2", err)
		return
	}
	//fmt.Println("读取到的buf=", this.Buf[:4])
	//	根据buf[:4]转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkglen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("发生丢包", err)
		return
	}

	//把pkg反序列化成 msg
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	//fmt.Println(mes)
	if err != nil {
		return
	}
	return

}

// 发送
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度给对方
	//此时 data就是发送消息|先发送 data的长度到服务器|
	//获取data长度，转为表示长度的切片
	//加密
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen) //长度转为 byte
	//defer conn.Close() //记得延时关闭|此处关闭会导致| use of closed network connection
	n, err := this.Conn.Write(this.Buf[0:4])
	if err != nil || n != 4 {
		fmt.Println("发送失败", err)
		return
	}
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("发送失败", err)
		return
	}
	return
}

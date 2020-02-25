package main

import (
	"chat/utils"
	"fmt"
	"os"
)

func main() {
	filePath := "D:/soft/GOPATH/src/chat/client/main/client.exe"
	size := utils.GetFileSize(filePath)
	open, _ := os.Open(filePath)
	sha1 := utils.FileSha1(open)
	fmt.Println(sha1)
	fmt.Println(size * 1.00 / 1024 / 1024)
}

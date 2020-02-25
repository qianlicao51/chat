package utils

import (
	"bufio"
	"fmt"
	"time"
)

func GetDateStr() (dateStr string) {
	now := time.Now()
	return now.Format("2006-01-02")
	//fmt.Println(now.Format("2006-01-02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02"))//年月日
	//fmt.Println(now.Format("15:04:05"))//时分秒
}

//格式化时间包含时分秒
func GetDateStrHmS() (dateStr string) {

	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
	//fmt.Println(now.Format("2006-01-02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//fmt.Println(now.Format("2006/01/02"))//年月日
	//fmt.Println(now.Format("15:04:05"))//时分秒
}

//打印空白行
func ShowBlackLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Println()
	}
}

//从控制台读取一行信息
func ReadLine(osIn *bufio.Reader) string {
	readLine, _, err := osIn.ReadLine()
	if err != nil {
		fmt.Println("从控制台读取一行信息发生错误", err)
		return ""
	}
	return string(readLine)
}

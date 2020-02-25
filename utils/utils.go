package utils

import (
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

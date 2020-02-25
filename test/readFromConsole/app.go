package main

import (
	"bufio"
	"fmt"
	"os"
)

//从控制台读取输入
func main() {
	var str string
	//length, err := fmt.Scanf("%s", &str)//注意使用%s读取输入字符串只能读取到空白符之前
	//fmt.Println(length,err,str)

	scanln, err := fmt.Scanln(&str)
	fmt.Println(scanln, err, str)

	in := bufio.NewReader(os.Stdin)
	line, prefix, err := in.ReadLine()
	fmt.Println(string(line), prefix, err)

}

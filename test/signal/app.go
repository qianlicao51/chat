package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {

	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//创建一个bool channel
	done := make(chan bool, 1)

	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	go func() {
		sig := <-sigs
		//go cmd https://blog.csdn.net/youngwhz1/article/details/88662172?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task
		//go shell https://studygolang.com/articles/7767
		exec.Command("cmd.exe", "/c", "start", "D:/pic/ps/20a01.jpg").Run() //没有start这个参数也行
		sprintln := fmt.Sprintln(sig.String(), sig.Signal)
		ioutil.WriteFile("d:/signal", []byte(sprintln), os.ModeAppend)
		done <- true
	}()

	//程序将在此处等待，直到它预期信号（如Goroutine所示）
	//在“done”上发送一个值，然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

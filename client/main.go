package main

import (
	"fmt"
	"os"
)

var (
	//	用户ID和密码
	userId  int
	userPwd string
)

func main() {

	//接受用户选择
	var key int
	//判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("--------------多人聊天系统--------")
		fmt.Println("\t\t\t 1 登陆")
		fmt.Println("\t\t\t 2 注册")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t （选择1-3）")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天系统")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			loop = false
		default:
			fmt.Println("重新输入")
		}
		// 根据用户输入，显示新的提示信息
		if key == 1 {
			//	要登陆
			fmt.Println("输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入用户密码")
			fmt.Scanf("%s\n", &userPwd)

			err := Login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败")
			} else {
				fmt.Println("登陆成功~~", err)
			}

		} else if key == 2 {
			fmt.Println("进行用户注册")
		}

	}

}

package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"reflect"
)

var Cfgs= ini.Empty()

func GetConf()  {
	err := ini.MapTo(Cfgs, "conf/conf.ini")
	Cfgs, err := ini.Load("conf/conf.ini")
	fmt.Println(Cfgs)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println("配置文件加载完毕:", Cfgs.Section("").Key("appName").String())
	fmt.Println(Cfgs.Section("client").Key("server_addr").String())
	fmt.Println(Cfgs.Section("server").Key("ser_port").String())
	fmt.Println(Cfgs.Section("redis").Key("redis_addr").String())
	fmt.Println("==",reflect.TypeOf(Cfgs))


}

func main() {
	GetConf()
	fmt.Println(Cfgs,"--")
	log.Println(Cfgs.Section("redis").Key("redis_addr").String())
}

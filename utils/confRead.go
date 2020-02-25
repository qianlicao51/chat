package utils

import (
	"fmt"
	"github.com/huandu/xstrings"
	"gopkg.in/ini.v1"
	"os"
)

var Cfg = ini.Empty()
//配置文件读取|https://ini.unknwon.io/docs/howto/load_data_sources
func init() {
	cfg, err := ini.LooseLoad("conf/conf.ini", "../../conf/conf.ini")
	if err != nil {
		fmt.Printf("Fail to read conf file: %v", err)
		os.Exit(1)
	}
	Cfg = cfg
	fmt.Println(xstrings.Center("配置文件加载完毕:", 30, "-"))

}

// 根据 selection key获取string Val
func ConfGetValString(section, key string) string {
	return Cfg.Section(section).Key(key).String()
}

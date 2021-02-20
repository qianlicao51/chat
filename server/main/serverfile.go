package main

import (
	"chat/common/iniconst"
	"chat/utils"
	"github.com/henrylee2cn/goutil"
	"os"
)

// initServerSaveFileDir 如果服务器保存 文件的 文件夹不存在创建文件夹
func initServerSaveFileDir() {
	var saveFileBaseDIR = utils.ConfGetValString(iniconst.INI_NAME_SERVER, iniconst.SERVER_SAVEFILE)
	if !goutil.FileExists(saveFileBaseDIR) {
		_ = os.MkdirAll(saveFileBaseDIR, os.ModePerm)
	}
}

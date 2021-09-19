package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	HttpName string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("read config file failed, err: %v", err)
	}
	loadServer(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	HttpName = file.Section("server").Key("HttpName").MustString("localhost")
}

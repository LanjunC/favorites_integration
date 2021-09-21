package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	HttpName   string
	PinboxAuth string
)

func InitSetting(isPrivacy bool) {
	var file *ini.File
	var err error
	if isPrivacy {
		file, err = ini.Load("./config/config_with_privacy.ini")
	} else {
		file, err = ini.Load("./config/config_without_privacy.ini")

	}
	if err != nil {
		fmt.Printf("read config file failed, err: %v", err)
	}
	loadServer(file)
	loadAuthorization(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	HttpName = file.Section("server").Key("HttpName").MustString("localhost")
}

func loadAuthorization(file *ini.File) {
	PinboxAuth = file.Section("authorization").Key("PinboxAuth").MustString("")
}

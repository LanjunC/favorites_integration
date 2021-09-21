package main

import (
	"codingcrea/favorites_integration/routes"
	"codingcrea/favorites_integration/utils"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

var (
	withPrivacy = flag.Bool("privacy", true, "use dev env config file...")
)

func main() {
	//处理os.args
	fmt.Printf("use privacy config file: %v.\n", *withPrivacy)
	flag.Parse()

	utils.InitSetting(*withPrivacy)

	defer glog.Flush()
	err := routes.InitRouter().Run(utils.HttpPort)
	if err != nil {
		fmt.Printf("start gin http server failed: %v", err)
	}
}

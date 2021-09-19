package main

import (
	"codingcrea/favorites_integration/routes"
	"codingcrea/favorites_integration/utils"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main()  {
	flag.Parse()
	defer glog.Flush()
	err := routes.InitRouter().Run(utils.HttpPort)
	if err != nil {
		fmt.Printf("start gin http server failed: %v", err)
	}
}



package main

import (
	"github.com/codinmoldovanu/mikrotik_auto_ssl/cli"
	"github.com/codinmoldovanu/mikrotik_auto_ssl/router"
)

func main() {
	routerInfo := cli.RequestRouterInfo()

	router.Assign(routerInfo)
	router.TestConnection()
	router.GetOutboundIP()
}

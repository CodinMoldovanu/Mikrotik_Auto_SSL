package main

import (
	"log"
	"os"

	"github.com/codinmoldovanu/mikrotik_auto_ssl/cli"
	"github.com/codinmoldovanu/mikrotik_auto_ssl/models"
	"github.com/codinmoldovanu/mikrotik_auto_ssl/router"
	"github.com/joho/godotenv"
)

func main() {
	routerInfo := models.RouterInfo{}
	err := godotenv.Load()
	if err != nil {
		routerInfo = cli.RequestRouterInfo()

	} else {
		log.Print(".env file found, asking for credentials")
		routerInfo.IPAddress = os.Getenv("IP")
		routerInfo.Port = os.Getenv("PORT")
		routerInfo.Username = os.Getenv("USERNAME")
		routerInfo.Password = os.Getenv("PASSWORD")
	}
	router.Assign(routerInfo)
	router.TestConnection()
	router.GetOutboundIP()
}

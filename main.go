package main

import (
	"log"
	"os"
	"os/exec"

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
		log.Print(".env file found, not asking for credentials")
		routerInfo.IPAddress = os.Getenv("IP")
		routerInfo.Port = os.Getenv("PORT")
		routerInfo.Username = os.Getenv("USERNAME")
		routerInfo.Password = os.Getenv("PASSWORD")
	}
	router.Assign(routerInfo)
	// router.TestConnection()
	if !router.NATRuleCheck() {
		router.CreateNAT()
	}

	cmd := exec.Command("/usr/bin/certbot", "certonly", "--manual")
	out, err := cmd.Output()
	if err == nil {
		println(string(out))
	} else {
		println(err.Error())
	}
}

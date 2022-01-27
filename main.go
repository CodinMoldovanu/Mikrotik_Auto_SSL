package main

import (
	"fmt"
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
	router.TestConnection()
	fmt.Print("Checking existing NAT rules or creating others...")
	if !router.NATRuleCheck() {
		fmt.Print("Creating NAT rule...")
		router.CreateNAT()
	}

	hostname := cli.RequestHostname()
	email := cli.RequestEmail()

	cert := exec.Command("sudo", "certbot", "certonly", "--standalone", "--preferred-challenges", "http", "-d", hostname, "--agree-tos", "-m", email)
	output, err := cert.CombinedOutput()
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Print(string(output))

	router.DisableNAT()

}

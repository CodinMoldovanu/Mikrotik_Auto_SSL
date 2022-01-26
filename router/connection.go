package router

import (
	"fmt"
	"log"
	"net"

	"github.com/codinmoldovanu/mikrotik_auto_ssl/models"
	"gopkg.in/routeros.v2"
)

var ri models.RouterInfo

func Assign(rif models.RouterInfo) {
	ri = rif
}

func dial() (*routeros.Client, error) {
	return routeros.Dial(ri.IPAddress+":"+ri.Port, ri.Username, ri.Password)
}

func TestConnection() bool {

	c, err := dial()

	if err != nil {
		fmt.Print(err.Error())
		log.Fatal("Connecting failed.")
	}
	defer c.Close()
	fmt.Print("Connection successfull.")
	return true
}

func NATRuleCheck() bool {
	exists := false
	c, err := dial()
	c.Async()
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal("Connecting failed.")
	}

	localIP := GetOutboundIP()

	r, err := c.Run("/ip/firewall/nat/print", "?chain=dstnat", "?to-ports=80", "?to-addresses="+localIP.String())
	if err != nil {
		fmt.Print(err.Error())
	}
	for _, rule := range r.Re {

		if rule.List[3].Key == "to-addresses" && rule.List[3].Value == localIP.String() && rule.List[len(rule.List)-2].Value == "true" {
			fmt.Println("A NAT rule to this place already exists, enabling it.")
			// args := []string{"/ip/firewall/nat/enable", "?number=" + rule.List[0].Value} #this doesn't work like that
			args := []string{"/ip/firewall/nat/enable", "+find[comment=\"AUTOSSL\"]"}
			e, err := c.RunArgs(args)
			if err != nil {
				log.Printf(err.Error())
			}
			log.Print(e.Done.String())
			exists = true
			break
		}
	}

	fmt.Print(r.Re)
	defer c.Close()
	log.Print(exists)
	return exists
}

func CreateNAT() error {
	localIP := GetOutboundIP()
	c, err := dial()
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal("Connecting failed.")
	}
	args := []string{"/ip/firewall/nat/add", "=chain=dstnat", "=action=dst-nat", "=to-addresses=" + localIP.String(), "=to-ports=80", "=protocol=tcp", "=in-interface=RDS_1GBps", "=dst-port=80", "=log=yes", "=comment=AUTOSSL"}

	n, err := c.RunArgs(args)
	if err != nil {
		log.Fatal(err.Error())
	}
	n.Done.String()
	defer c.Close()

	return err
}

// Solution stolen from Marcel Molina @ https://stackoverflow.com/a/37382208
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Print(localAddr.IP)
	return localAddr.IP
}

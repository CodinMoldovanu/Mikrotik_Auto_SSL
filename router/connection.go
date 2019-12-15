package router

import (
	"fmt"
	"log"
	"net"
	"strings"

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

	localIP := GetOutboundIP()

	// r, err := c.Run("/ip/firewall/nat/ print to-ports=\"80\"?disabled=true")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	r, err := c.RunArgs(strings.Split("/ip/firewall/nat/print?to-ports=\"80\"?disabled=false", "?"))
	if err != nil {
		fmt.Print(err.Error())
	}

	for _, rule := range r.Re {
		fmt.Print(rule.String())
		if rule.List[3].Value == localIP.String() {
			fmt.Print("A NAT rule to this place already exists.")
			fmt.Print(c.Run("/ip/firewall/nat", "?enable="+rule.List[0].Value))
			break
		} else {
			fmt.Print("Creating a new NAT rule.")
			n, err := c.Run("/ip/firewall/nat/print", "?add", "?chain=dstnat", "?action=dst-nat", "?to-addresses="+localIP.String(), "?to-ports=80", "?protocol=tcp", "?in-interface=RDS_1Gbps", "?dst-port=80", "?log=yes")
			fmt.Print(n)
			if err != nil {
				log.Fatal(n)
			}
			break
		}
	}

	if r != nil {
		fmt.Print("Connection to router is successful.")
		return true
	}
	return false
}

func NATRuleCheck() bool {
	c, err := dial()
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal("Connecting failed.")
	}
	defer c.Close()

	localIP := GetOutboundIP()

	// r, err := c.Run("/ip/firewall/nat/ print to-ports=\"80\"?disabled=true")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	r, err := c.RunArgs(strings.Split("/ip/firewall/nat/print?to-ports=\"80\"?to-addresses="+localIP.To16().String()+"?disabled=false", "?"))
	if err != nil {
		fmt.Print(err.Error())
	}

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

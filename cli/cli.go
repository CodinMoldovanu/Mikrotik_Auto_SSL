package cli

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/codinmoldovanu/mikrotik_auto_ssl/models"
	"github.com/manifoldco/promptui"
)

// func RequestTOS() bool {
// 	promptTOS := promptui.Prompt{
// 		Label: "Do you accept the TOS of the CA? Y/n",
// 	}
// 	// if promptTOS == "Y" {
// 		return true
// 	}
// 	return false
// }

func RequestRouterInfo() models.RouterInfo {

	validateIPv4 := func(ip string) error {
		ipBlockCount := strings.Split(ip, ".")

		if len(ipBlockCount) != 4 {
			return errors.New("The string you provided isn't a valid IPv4 address.")
		}

		for _, x := range ipBlockCount {
			if i, err := strconv.Atoi(x); err == nil {
				if i > 0 || i < 255 {
					return nil
				} else {
					return errors.New("The string you provided is out of the IPv4 bounds.")
				}
			}
		}
		return nil
	}

	validatePort := func(port string) error {
		if _, err := strconv.Atoi(port); err == nil {
			return nil
		} else {
			return errors.New("Not a valid port.")
		}
	}

	promptIP := promptui.Prompt{
		Label:    "Enter the IPv4 address of the Mikrotik router:",
		Validate: validateIPv4,
	}

	promptPort := promptui.Prompt{
		Label:    ("Enter the port of the Mikrotik router:"),
		Validate: validatePort,
	}

	promptUsername := promptui.Prompt{
		Label: "Enter the username for the Mikrotik router:",
	}

	promptPassword := promptui.Prompt{
		Label: ("Enter the password for the user you provided"),
		Mask:  '*',
	}

	promptWANin := promptui.Prompt{
		Label: ("Enter the WAN In Interface name:"),
	}

	ipAddress, err := promptIP.Run()

	port, err := promptPort.Run()

	username, err := promptUsername.Run()

	password, err := promptPassword.Run()

	wanIN, err := promptWANin.Run()

	if err != nil {
		fmt.Printf("something went wrong %v", err)
	}

	var routerInf = models.RouterInfo{}
	routerInf.IPAddress = ipAddress
	routerInf.Port = port
	routerInf.Username = username
	routerInf.Password = password
	routerInf.WANIn = wanIN
	return routerInf
}

func RequestHostname() string {
	promptDomain := promptui.Prompt{
		Label: ("Enter the domain for which the SSL cert will be generated:"),
	}

	domain, err := promptDomain.Run()
	if err != nil {
		log.Fatal("Error at hostname input")
	}

	return domain
}

func RequestEmail() string {
	promptEmail := promptui.Prompt{
		Label: ("Enter your email account for important account notifications:"),
	}

	email, err := promptEmail.Run()
	if err != nil {
		log.Fatal("Error at email input")
	}
	return email
}

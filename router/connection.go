package router

import (
	"fmt"
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
	}
	defer c.Close()

	r, err := c.RunArgs(strings.Split("/system/resource/print", " "))
	if err != nil {
		fmt.Print(err.Error())
	}

	// fmt.Print(r)

	if r != nil {
		fmt.Print("Connection to router is successful.")
		return true
	}
	return false
}

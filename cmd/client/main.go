package main

import (
	"fmt"
	"time"

	"github.com/lopygo/lopy_ddns/ip/ipv6/test_ipv6"

	"github.com/lopygo/lopy_ddns/ip/ipv4/sohu"

	"github.com/lopygo/lopy_ddns/ip/common"
	"github.com/lopygo/lopy_ddns/model/current"
	currentService "github.com/lopygo/lopy_ddns/service/current"
)

func main() {
	fmt.Println("client simple")

	// current data
	currentModel := currentService.CurrentModel()
	currentModel.SetOnIpv4Changed(ipv4Changed)
	currentModel.SetOnIpv6Changed(ipv6Changed)

	// init address of ipv4 and ipv6
	// default use sohu in ipv4
	go resolveIP(&sohu.IpDriver{})
	// default use test_ipv6 in ipv6
	go resolveIP(&test_ipv6.IpDriver{})

	// loop get address

	// default use sohu in ipv6

	//

	time.Sleep(10 * time.Second)
}

func resolveIP(i common.IDriver) {
	ip, err := i.Resolve()
	if err != nil {
		return
	}

	if len(ip) == 0 {
		return
	}

	// set ip
	if common.IsIpv4(ip) {
		currentService.CurrentModel().SetIPV4(ip)
	} else if common.IsIpv6(ip) {
		currentService.CurrentModel().SetIPV6(ip)
	}
}

func ipv4Changed(model *current.CurrentModel, newIP string) {
	fmt.Println("ipv4 changed: ", newIP)
}

func ipv6Changed(model *current.CurrentModel, newIP string) {
	fmt.Println("ipv6 changed: ", newIP)
}

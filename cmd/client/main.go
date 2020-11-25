package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lopygo/lopy_ddns/ip/ipv6/test_ipv6"

	"github.com/lopygo/lopy_ddns/ip/ipv4/sohu"

	"github.com/lopygo/lopy_ddns/ip/common"
	"github.com/lopygo/lopy_ddns/model/current"
	currentService "github.com/lopygo/lopy_ddns/service/current"
	"github.com/lopygo/lopy_ddns/service/driver_list"
)

var loopLocker sync.Mutex
var started bool

func main() {
	fmt.Println("client simple")

	// current data
	currentModel := currentService.CurrentModelService()
	currentModel.SetOnIpv4Changed(ipv4Changed)
	currentModel.SetOnIpv6Changed(ipv6Changed)

	// init address of ipv4 and ipv6
	// default use sohu in ipv4
	go resolveIP(&sohu.IpDriver{})
	// default use test_ipv6 in ipv6
	go resolveIP(&test_ipv6.IpDriver{})

	//

	// loop get address

	// driver list

	driverList := driver_list.ListService()
	for _, v := range driverList.Ipv4All() {
		fmt.Println("print ipv4 drivers")
		fmt.Println(v)
	}

	for _, v := range driverList.Ipv6All() {
		fmt.Println("print ipv6 drivers")
		fmt.Println(v)
	}

	// default use sohu in ipv6

	//

	go loop()

	time.Sleep(1000 * time.Second)
	started = false
}

func loop() {

	if started {
		return
	}

	loopLocker.Lock()
	defer loopLocker.Unlock()
	started = true
	ticker := time.NewTicker(time.Duration(300) * time.Second)
	for started {
		select {
		case <-ticker.C:
			fmt.Println("check ip")
		}
	}
	ticker.Stop()

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
		currentService.CurrentModelService().SetIPV4(ip)
	} else if common.IsIpv6(ip) {
		currentService.CurrentModelService().SetIPV6(ip)
	}
}

func ipv4Changed(model *current.CurrentModel, newIP string) {
	fmt.Println("ipv4 changed: ", newIP)
	// driver_list.ListService().UpdateIpv4(newIP)
}

func ipv6Changed(model *current.CurrentModel, newIP string) {
	fmt.Println("ipv6 changed: ", newIP)
	// driver_list.ListService().UpdateIpv6(newIP)
}

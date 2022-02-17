package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/lopygo/lopy_ddns/ip/ipv6/test_ipv6"

	"github.com/lopygo/lopy_ddns/ip/ipv4/sohu"

	"github.com/lopygo/lopy_ddns/ip/common"
	"github.com/lopygo/lopy_ddns/model/current"
	"github.com/lopygo/lopy_ddns/service/about"
	currentService "github.com/lopygo/lopy_ddns/service/current"
	"github.com/lopygo/lopy_ddns/service/driver_list"
)

var loopLocker sync.Mutex

func main() {

	aboutModel, err := about.FromInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	showV := flag.Bool("v", false, "version")
	flag.Parse()
	if *showV {
		l := make([]string, 0)
		l = append(l, fmt.Sprintf("%s\t%s", aboutModel.AppName, aboutModel.AppVersion))
		l = append(l, "\n")
		l = append(l, fmt.Sprintf("Built Time:\t%s", aboutModel.BuildTime))
		l = append(l, fmt.Sprintf("Git Commit:\t%s", aboutModel.GITCommit))
		l = append(l, fmt.Sprintf("Built Go:\t%s", aboutModel.BuildGoVersion))
		l = append(l, fmt.Sprintf("Website:\t%s", aboutModel.WebSite))

		fmt.Println(strings.Join(l, "\n"))
		return
	}

	fmt.Println("client simple")
	ctx, cancel := context.WithCancel(context.Background())

	// current data
	currentModel := currentService.CurrentModelService()
	currentModel.SetOnIpv4Changed(ipv4Changed)
	currentModel.SetOnIpv6Changed(ipv6Changed)

	//

	// loop get address

	// driver list

	driverList := driver_list.ListService()
	for _, v := range driverList.Ipv4All() {
		fmt.Println("print ipv4 drivers")
		fmt.Printf("%++v \n", v)
		err := v.Init()
		if err != nil {
			fmt.Println("a driver init error: ", err)
		}
	}

	for _, v := range driverList.Ipv6All() {
		fmt.Println("print ipv6 drivers")
		fmt.Printf("%++v \n", v)
		err := v.Init()
		if err != nil {
			fmt.Println("a driver init error", err)
		}
	}

	// default use sohu in ipv6

	//

	go loop(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		getPublicIp()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:

			cancel()
			return
		case <-ctx.Done():

			return
		}
	}
}

func loop(ctx context.Context) {

	loopLocker.Lock()
	defer loopLocker.Unlock()

	ticker := time.NewTicker(time.Duration(300) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("check ip")
			getPublicIp()
		case <-ctx.Done():
			return
		}
	}

}

func getPublicIp() {

	// init address of ipv4 and ipv6
	// default use sohu in ipv4
	go getCurrentIP(&sohu.IpDriver{})
	// default use test_ipv6 in ipv6
	go getCurrentIP(&test_ipv6.IpDriver{})

}

// getCurrentIP get ip
func getCurrentIP(i common.IDriver) {
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

func ipv4Changed(model *current.CurrentModel, newLocalIP string) {
	fmt.Println("ipv4 changed: ", newLocalIP, len(newLocalIP))
	// driver_list.ListService().UpdateIpv4(newIP)

	fmt.Println("get ipv4 from ddns driver")

	driverList := driver_list.ListService()
	for _, v := range driverList.Ipv4All() {
		fmt.Println("print ipv4 drivers")
		fmt.Printf("%++v \n", v)

		ddnsIP, err := v.ResolveIP()
		if err != nil {
			fmt.Printf("lookup %s err: %v \n", v.Host(), err)
			continue
		}

		if ddnsIP == newLocalIP {
			fmt.Println("local ip equal resolve ip")
			continue
		}
		//
		fmt.Println()
		fmt.Printf("lookup ip:   %s : %s \n", v.Host(), ddnsIP)
		fmt.Println("do change ipv4 for ddns where ip changed")

		err = v.Update(newLocalIP)
		fmt.Println("update ip err: ", err)
	}

}

func ipv6Changed(model *current.CurrentModel, newLocalIP string) {
	fmt.Println("ipv6 changed: ", newLocalIP)
	// driver_list.ListService().UpdateIpv6(newIP)
	fmt.Println("get ipv6 from ddns driver")
	fmt.Println("do change ipv6 for ddns where ip changed")
}

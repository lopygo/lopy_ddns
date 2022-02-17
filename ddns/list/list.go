package list

import (
	"fmt"

	"github.com/lopygo/lopy_ddns/config"
	driver2 "github.com/lopygo/lopy_ddns/ddns/driver"
	"github.com/lopygo/lopy_ddns/ddns/driver/driver"
	"github.com/lopygo/lopy_ddns/ip/common"
)

type DriverList struct {
	listIpv4 []driver.IDriver
	listIpv6 []driver.IDriver
}

func (p *DriverList) Ipv4All() []driver.IDriver {
	return p.listIpv4
}

func (p *DriverList) Ipv6All() []driver.IDriver {
	return p.listIpv6
}

func (p *DriverList) UpdateIpv4(ip string) {
	if !common.IsIpv4(ip) {
		return
	}
	for _, v := range p.Ipv4All() {
		err := UpdateDDNS(v, ip)

		if err != nil {
			fmt.Printf("update error: %s, %s, %+v", v.Host(), ip, err)
		} else {
			fmt.Printf("update success: %s, %s", v.Host(), ip)
		}
	}
}

func (p *DriverList) UpdateIpv6(ip string) {
	if !common.IsIpv6(ip) {
		return
	}
	for _, v := range p.Ipv6All() {
		err := UpdateDDNS(v, ip)
		if err != nil {
			fmt.Printf("update error: %s, %s, %+v\n", v.Host(), ip, err)
		} else {
			fmt.Printf("update success: %s, %s\n", v.Host(), ip)
		}
	}
}

func UpdateDDNS(driverInstance driver.IDriver, ip string) error {

	oldIp, err := driverInstance.ResolveIP()
	if err != nil {
		return err
	}

	// 目前取0 ？还是怎么办
	if len(oldIp) == 0 {
		return fmt.Errorf("no ip resolved on [%s]", driverInstance.Host())
	}

	// 不做，不管
	if oldIp == ip {
		return fmt.Errorf("ip is same, ignored")
	}

	err = driverInstance.UpdateBefore()
	if err != nil {
		return err
	}

	err = driverInstance.Update(ip)

	if err != nil {
		return err
	}

	return driverInstance.UpdateAfter()
}

func LoadFromDriversConfig(confList []config.Driver) (*DriverList, error) {

	list := new(DriverList)
	list.listIpv4 = make([]driver.IDriver, 0)
	list.listIpv6 = make([]driver.IDriver, 0)

	listAvail := driver2.ListAvailable()

	for _, v := range confList {

		creator, has := listAvail[v.Driver]

		if !has {
			return nil, fmt.Errorf("dirver [%s] can not be supposed", v.Driver)
		}

		f, err := creator(v)
		if err != nil {
			return nil, err
		}

		if v.GetType() == config.IPTypeV6 {
			list.listIpv6 = append(list.listIpv6, f)
		} else {
			list.listIpv4 = append(list.listIpv4, f)
		}
	}

	return list, nil
}

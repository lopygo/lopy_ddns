package driver

import (
	"github.com/lopygo/lopy_ddns/config"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod"
	driver2 "github.com/lopygo/lopy_ddns/ddns/driver/driver"
)

type DriverCreator func(config.Driver) (driver2.IDriver, error)

func ListAvailable() map[string]DriverCreator {

	l := make(map[string]DriverCreator, 0)
	l["dnspod"] = dnspod.LoadFromDriverConfig
	l["dnspod.cn"] = dnspod.LoadFromDriverConfig
	return l
}

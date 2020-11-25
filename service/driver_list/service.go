package driver_list

import (
	"fmt"
	"sync"

	"github.com/lopygo/lopy_ddns/ddns/list"

	"github.com/lopygo/lopy_ddns/service/config"
)

var _list *list.DriverList
var _once sync.Once

func ListService() *list.DriverList {
	_once.Do(func() {
		var err error
		conf := config.ConfigService()
		_list, err = list.LoadFromDriversConfig(conf.Drivers)
		if err != nil {
			panic(fmt.Sprintf("init drivers error: %+v", err))
		}
	})
	return _list
}

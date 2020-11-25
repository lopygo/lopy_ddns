package config

import (
	"fmt"
	"sync"

	config2 "github.com/lopygo/lopy_ddns/config"
)

var _config config2.Config
var _once sync.Once

func ConfigService() config2.Config {
	_once.Do(func() {
		var err error

		_config, err = config2.LoadFromFile()
		if err != nil {
			panic(fmt.Sprintf("load config error: %+v", err))
		}
	})
	return _config
}

package current

import (
	"sync"

	current2 "github.com/lopygo/lopy_ddns/model/current"
)

var _currentData *current2.CurrentModel

var _once sync.Once

func CurrentModelService() *current2.CurrentModel {
	_once.Do(func() {
		_currentData = current2.NewCurrentModel()
	})

	return _currentData
}

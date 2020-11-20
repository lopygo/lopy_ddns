package common

import (
	"sync"
)

type IAdapter interface {
	Check() error

	GetData() map[string]string

	Method() string
}

type AAdapter struct {
	data map[string]string

	once sync.Once

	locker sync.Mutex
}

func (p *AAdapter) GetData() map[string]string {
	return p.list()
}

func (p *AAdapter) list() map[string]string {
	p.once.Do(func() {
		p.data = make(map[string]string)
	})
	return p.data
}

func (p *AAdapter) Get(key string) string {
	v, has := p.data[key]
	if !has {
		return ""
	}

	return v
}

func (p *AAdapter) Set(key string, value string) {
	p.locker.Lock()
	m := p.list()
	m[key] = value
	// p.data[key] = value
	p.locker.Unlock()
}

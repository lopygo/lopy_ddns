package record

import (
	"fmt"
	"strings"
	"sync"
)

type domainField struct {
	domain   string
	domainID int
}

func (p *domainField) SetDomain(domain string) {
	p.domain = domain
}

func (p *domainField) SetDomainID(domainID int) {
	p.domainID = domainID
}

func (p *domainField) check() error {

	if p.domainID < 0 {
		p.domainID = 0
	}

	p.domain = strings.Trim(p.domain, " ")

	//

	if p.domainID > 0 {
		return nil
	}

	if len(p.domain) > 0 {
		return nil
	}

	return fmt.Errorf("%s 不能同时为空", "DomainID 和 Domain ")
}

type recordLineField struct {
	recordLine   string
	recordLineID int
}

func (p *recordLineField) SetRecordLine(recordLine string) {
	p.recordLine = recordLine
}

func (p *recordLineField) SetRecordLineID(recordLineID int) {
	p.recordLineID = recordLineID

}

func (p *recordLineField) check() error {

	if p.recordLineID < 0 {
		p.recordLineID = 0
	}

	p.recordLine = strings.Trim(p.recordLine, " ")

	//

	if p.recordLineID > 0 {
		return nil
	}

	if len(p.recordLine) == 0 {
		p.recordLine = "默认"
	}

	return nil
}

type adapterAbstract struct {
	data map[string]string

	once sync.Once

	locker sync.Mutex
}

func (p *adapterAbstract) GetPostData() map[string]string {
	return p.list()
}

func (p *adapterAbstract) list() map[string]string {
	p.once.Do(func() {
		p.data = make(map[string]string)
	})
	return p.data
}

func (p *adapterAbstract) Get(key string) string {
	v, has := p.data[key]
	if !has {
		return ""
	}

	return v
}

func (p *adapterAbstract) Set(key string, value string) {
	p.locker.Lock()
	m := p.list()
	m[key] = value
	// p.data[key] = value
	p.locker.Unlock()
}

type IAdapter interface {
	Check() error

	GetPostData() map[string]string

	Method() string

	SetResponseJson(jsonBuf []byte) (err error)
}

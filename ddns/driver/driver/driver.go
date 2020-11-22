package driver

import (
	"net"
)

// IDriver
type IDriver interface {
	UpdateBefore() error

	Update(string) error

	UpdateAfter() error

	ResolveIp() (ip []string, err error)
}

// ADriver
type ADriver struct {
	lookupHost string
}

// SetHost
func (p *ADriver) SetHost(host string) {
	p.lookupHost = host
}

func (p *ADriver) Host() string {
	return p.lookupHost
}

func (p *ADriver) ResolveIp() ([]string, error) {
	// m := IPModel{}
	addr, err := net.LookupIP(p.Host())
	if err != nil {
		return nil, err
	}

	l := make([]string, 0)
	// 暂时不考虑ipv6
	for _, v := range addr {
		l = append(l, v.String())
	}

	return l, nil
}

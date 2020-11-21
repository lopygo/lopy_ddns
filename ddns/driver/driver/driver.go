package driver

import (
	"net"
)

// IDriver
type IDriver interface {
	UpdateBefore() error

	Update() error

	UpdateAfter() error

	ResolveIp() (IPModel, err error)
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

func (p *ADriver) ResolveIp() (IPModel, error) {
	m := IPModel{}
	addr, err := net.LookupIP(p.Host())
	if err != nil {
		return m, err
	}

	l := make([]string, 0)
	// 暂时不考虑ipv6
	for _, v := range addr {
		l = append(l, v.String())
		m.IPV4 = v.String()
	}

	return m, nil
}

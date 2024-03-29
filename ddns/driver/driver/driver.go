package driver

import (
	"net"
	"time"
)

// IDriver
type IDriver interface {
	Init() error

	UpdateBefore() error

	Update(string) error

	UpdateAfter() error

	ResolveIP() (ip string, err error)

	Host() string

	LastIP() string

	LastUpdateTime() time.Time
}

// ADriver
type ADriver struct {
	lookupHost string

	lastIP string

	lastUpdateTime time.Time
}

func (p *ADriver) LastIP() string {
	return p.lastIP
}

func (p *ADriver) SetLastIP(ip string) {
	p.lastIP = ip
}

// SetHost
func (p *ADriver) SetHost(host string) {
	p.lookupHost = host
	p.lastUpdateTime = time.Now()
}

func (p *ADriver) Host() string {
	return p.lookupHost
}

func (p *ADriver) LastUpdateTime() time.Time {
	return p.lastUpdateTime
}

func (p *ADriver) ResolveIP() (string, error) {
	// m := IPModel{}
	addr, err := net.LookupIP(p.Host())
	if err != nil {
		return "", err
	}

	// 只取第一个
	ip := addr[0].String()

	p.SetLastIP(ip)
	return ip, nil
}

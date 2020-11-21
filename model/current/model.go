package current

import "sync"

type OnIpv4Changed func(sender *CurrentModel, newIp string)
type OnIpv6Changed func(sender *CurrentModel, newIp string)

type CurrentModel struct {
	// IPV4 string
	// IPV6 string
	ipv4 string
	ipv6 string

	lockerIpv4 sync.Mutex
	lockerIpv6 sync.Mutex

	onIpv4Changed OnIpv4Changed
	onIpv6Changed OnIpv6Changed
}

func NewCurrentModel() *CurrentModel {
	i := new(CurrentModel)

	return i
}
func (p *CurrentModel) IPV4() string {
	return p.ipv4
}
func (p *CurrentModel) SetIPV4(ip string) {
	p.lockerIpv4.Lock()
	oldIP := p.ipv4
	p.ipv4 = ip
	if oldIP != ip && p.onIpv4Changed != nil {
		go p.onIpv4Changed(p, ip)
	}
	p.lockerIpv4.Unlock()
}

func (p *CurrentModel) IPV6() string {
	return p.ipv6
}
func (p *CurrentModel) SetIPV6(ip string) {
	p.lockerIpv6.Lock()
	oldIP := p.ipv6
	p.ipv6 = ip
	if oldIP != ip && p.onIpv6Changed != nil {
		go p.onIpv6Changed(p, ip)
	}
	p.lockerIpv6.Unlock()
}

func (p *CurrentModel) SetOnIpv4Changed(callback OnIpv4Changed) {
	p.onIpv4Changed = callback
}

func (p *CurrentModel) SetOnIpv6Changed(callback OnIpv6Changed) {
	p.onIpv6Changed = callback
}

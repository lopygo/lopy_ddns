package ipip_net

import (
	"fmt"
	"regexp"

	"github.com/lopygo/lopy_ddns/ip/common"
)

var _ common.IDriver = new(IpDriver)

const resolveUrl = "https://myip.ipip.net"

type IpDriver struct {
}

func (p *IpDriver) Resolve() (string, error) {

	buf, err := common.DoRequest(resolveUrl)
	if err != nil {
		return "", err
	}
	m, err := modelFromBuffer(buf)
	if err != nil {
		return "", err
	}

	return m.IP, nil
}

// `
// 当前 IP：123.145.118.132  来自于：中国 重庆 重庆  联通
// `

type ipModel struct {
	IP string
}

func modelFromBuffer(buf []byte) (ipModel, error) {
	m := ipModel{}

	s := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	res := s.FindSubmatch(buf)

	if len(res) == 0 {
		return m, fmt.Errorf("there is no ip mathed")
	}

	m.IP = string(res[1])

	return m, nil
}

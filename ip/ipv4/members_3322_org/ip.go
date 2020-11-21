package members_3322_org

import (
	"fmt"
	"regexp"

	"github.com/lopygo/lopy_ddns/ip/common"
)

var _ common.IDriver = new(IpDriver)

const resolveUrl = "http://members.3322.org/dyndns/getip"

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
// IP	: 123.145.118.132
// 地址	: 中国  重庆
// 运营商	: 联通

// 数据二	: 重庆市 | 联通

// 数据三	:

// URL	: http://www.cip.cc/123.145.118.132

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

package ifconfig_me

import "github.com/lopygo/lopy_ddns/ip/common"

const resolveUrl = "https://ifconfig.me"

var _ common.IDriver = new(IpDriver)

type IpDriver struct {
}

func (p *IpDriver) Resolve() (string, error) {

	buf, err := common.DoRequest(resolveUrl)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

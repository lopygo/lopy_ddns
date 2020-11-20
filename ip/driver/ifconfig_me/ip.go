package ifconfig_me

import "github.com/lopygo/lopy_ddns/ip/driver/common"

const resolveUrl = "https://ifconfig.me"

type IpDriver struct {
}

func (p *IpDriver) Resolve() (string, error) {

	buf, err := common.DoRequest(resolveUrl)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

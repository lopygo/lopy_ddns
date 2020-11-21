package ipinfo_io

import (
	"encoding/json"

	"github.com/lopygo/lopy_ddns/ip/common"
)

var _ common.IDriver = new(IpDriver)

const resolveUrl = "https://ipinfo.io"

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
// {
// 	"ip": "123.145.118.132",
// 	"city": "Chongqing",
// 	"region": "Chongqing",
// 	"country": "CN",
// 	"loc": "29.5603,106.5577",
// 	"org": "AS4837 CHINA UNICOM China169 Backbone",
// 	"timezone": "Asia/Shanghai",
// 	"readme": "https://ipinfo.io/missingauth"
// }
// `

type ipModel struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func modelFromBuffer(buf []byte) (ipModel, error) {
	m := ipModel{}
	err := json.Unmarshal(buf, &m)
	return m, err
}

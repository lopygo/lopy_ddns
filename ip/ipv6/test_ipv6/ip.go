package test_ipv6

import (
	"encoding/json"

	"github.com/lopygo/lopy_ddns/ip/common"
)

var _ common.IDriver = new(IpDriver)

const resolveUrl = "http://ipv6.lookup.test-ipv6.com/ip/"

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
// 	"ip": "2408:8262:188c:544c:c464:d7:be38:78ef",
// 	"type": "ipv6",
// 	"subtype": "",
// 	"via": "",
// 	"padding": "",
// 	"asn": "4837",
// 	"asnlist": "4837",
// 	"asn_name": "CHINA169-BACKBONE CHINA UNICOM China169 Backbone",
// 	"country": "CN",
// 	"protocol": "HTTP/2.0"
// }
// `

type ipModel struct {
	IP      string `json:"ip"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Country string `json:"country"`
}

func modelFromBuffer(buf []byte) (ipModel, error) {
	m := ipModel{}
	err := json.Unmarshal(buf, &m)
	return m, err
}

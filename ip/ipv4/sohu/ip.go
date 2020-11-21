package sohu

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/lopygo/lopy_ddns/ip/common"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var _ common.IDriver = new(IpDriver)

const resolveUrl = "https://pv.sohu.com/cityjson"

type IpDriver struct {
}

func (p *IpDriver) Resolve() (string, error) {
	buf, err := common.DoRequest(resolveUrl)
	if err != nil {
		return "", err
	}

	reader := transform.NewReader(bytes.NewReader(buf), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	start := bytes.IndexByte(d, byte('{'))
	end := bytes.IndexByte(d, byte('}'))

	m, err := modelFromBuffer(d[start : end+1])
	if err != nil {
		return "", err
	}

	return m.IP, nil
}

type ipModel struct {
	IP   string `json:"cip"`
	ID   string `json:"cid"`
	Name string `json:"cname"`
}

func modelFromBuffer(buf []byte) (ipModel, error) {
	m := ipModel{}
	err := json.Unmarshal(buf, &m)
	return m, err
}

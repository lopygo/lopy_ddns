package dnspod

import (
	"fmt"
	"strings"

	"github.com/lopygo/lopy_ddns/config"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/adapter/record"
	configClient "github.com/lopygo/lopy_ddns/ddns/driver/dnspod/config"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/request"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/response"
	"github.com/lopygo/lopy_ddns/ddns/driver/driver"
	"github.com/mitchellh/mapstructure"
)

var _ driver.IDriver = new(Client)

type Client struct {
	driver.ADriver

	config config.Driver

	extConfig ExtConfig

	request *request.HttpClient

	dataDomain response.DomainInfo

	dataRecord response.RecordInfo
}

// UpdateBefore 这一步是为了取得 recordId 和 domain id
func (p *Client) UpdateBefore() error {

	adap := &record.ListAdapter{
		Domain: p.config.Domain,
	}

	buf, err := p.request.Request(adap)
	if err != nil {
		return err
	}

	res, err := response.ListResultFromBuffer(buf)

	if err != nil {
		return err
	}
	p.dataDomain = res.Domain

	// sub main
	line := strings.Trim(p.extConfig.Line, " ")
	if len(line) == 0 {
		line = "默认"
	}
	for _, v := range res.Records {
		if p.config.SubDomain == v.Name && v.Line == line {
			p.dataRecord = v
			break
		}
	}

	if p.dataRecord == (response.RecordInfo{}) {
		return fmt.Errorf(
			"the is no this record {sub_domain: \"%s\" , line: \"%s\"}",
			p.config.SubDomain,
			line,
		)
	}

	return nil
}

func (p *Client) Update(ip string) error {

	adap := &record.ModifyAdapter{
		DomainID:     string(p.dataDomain.ID),
		RecordID:     string(p.dataRecord.ID),
		SubDomain:    p.config.SubDomain,
		RecordLine:   p.dataRecord.Line,
		RecordLineID: p.dataRecord.LineID,
	}

	adap.SetValue(ip, p.config.GetType())

	buf, err := p.request.Request(adap)
	if err != nil {
		return err
	}

	modifyRes, err := response.ModifyResultFromBuffer(buf)
	if err != nil {
		return err
	}
	if modifyRes.Status.Code != "1" {
		return fmt.Errorf("update ip error: %s", modifyRes.Status.Message)
	}
	fmt.Println(modifyRes)

	return nil
}

func (p *Client) UpdateAfter() error {
	// panic("not implemented") // TODO: Implement
	return nil
}

func LoadFromDriverConfig(conf config.Driver) (driver.IDriver, error) {

	c := new(Client)
	c.config = conf
	clientConf := configClient.NewConfig(conf.Username, conf.Password)
	client := request.NewHttpClient(&clientConf)
	c.request = client
	e, err := ExtConfigLoad(conf.Ext)
	if err != nil {
		return nil, err
	}
	c.extConfig = e

	c.SetHost(fmt.Sprintf("%s.%s", c.config.SubDomain, c.config.Domain))
	return c, nil
}

type ExtConfig struct {
	Line  string
	Email string
}

func ExtConfigDeafult() ExtConfig {
	return ExtConfig{
		Line:  "默认",
		Email: "",
	}
}

func ExtConfigLoad(conf map[string]interface{}) (ExtConfig, error) {
	c := ExtConfigDeafult()
	err := mapstructure.Decode(conf, &c)
	return c, err
}

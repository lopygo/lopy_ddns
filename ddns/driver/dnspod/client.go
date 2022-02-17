package dnspod

import (
	"fmt"
	"strings"

	"github.com/lopygo/lopy_ddns/config"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/adapter/record"
	configClient "github.com/lopygo/lopy_ddns/ddns/driver/dnspod/config"
	"github.com/lopygo/lopy_ddns/ddns/driver/dnspod/request"
	"github.com/lopygo/lopy_ddns/ddns/driver/driver"
	"github.com/mitchellh/mapstructure"
)

var _ driver.IDriver = new(Client)

type Client struct {
	driver.ADriver

	config config.Driver

	extConfig ExtConfig

	// request *request.HttpClient

	dataDomain *Domain

	dataRecord *Record

	requestClient func() *request.HttpClient
}

// func (p *Client) requestClient() *request.HttpClient {

// }

// UpdateBefore 这一步是为了取得 recordId 和 domain id
func (p *Client) UpdateBefore() error {

	return nil
}

func (p *Client) Update(ip string) (err error) {
	if p.dataDomain == nil {
		err = fmt.Errorf("domain 不能为空")
		return
	}

	if p.dataRecord == nil {
		err = fmt.Errorf("record 不能为空")
		return
	}

	adap := &record.DdnsAdapter{
		RecordID:  p.dataRecord.ID,
		SubDomain: p.config.SubDomain,
	}
	adap.SetDomainID(p.dataDomain.ID)
	adap.SetValue(ip)

	_, err = p.requestClient().Request(adap)
	if err != nil {
		return err
	}

	modifyRes := adap.Response()
	fmt.Println(modifyRes)

	return nil
}

func (p *Client) UpdateAfter() error {
	// panic("not implemented") // TODO: Implement
	return nil
}

func (p *Client) ResolveIP() (ip string, err error) {
	if p.dataDomain == nil {
		err = fmt.Errorf("domain 不能为空")
		return
	}

	if p.dataRecord == nil {
		err = fmt.Errorf("record 不能为空")
		return
	}

	l := record.InfoAdapter{}
	l.SetDomainID(p.dataDomain.ID)
	l.RecordID = p.dataRecord.ID

	_, err = p.requestClient().Request(&l)
	if err != nil {
		return
	}

	ip = l.Response().Record.Value
	return
}

func (p *Client) Init() error {
	adap := &record.ListAdapter{}

	adap.SetDomain(p.config.Domain)
	adap.SubDomain = p.config.SubDomain
	_, err := p.requestClient().Request(adap)
	if err != nil {
		return err
	}

	// domain
	res := adap.Response()
	domainId, err := res.Domain.ID.Int64()
	if err != nil {
		return err
	}
	p.dataDomain = &Domain{
		ID:     int(domainId),
		Domain: res.Domain.Domain,
	}

	//

	// sub main
	line := strings.Trim(p.extConfig.Line, " ")
	if len(line) == 0 {
		line = "默认"
	}

	//
	for _, v := range res.Records {
		if p.config.SubDomain == v.SubDomain && v.Line == line {

			recordId, err := v.ID.Int64()
			if err != nil {
				continue
			}
			p.dataRecord = &Record{
				ID:        int(recordId),
				SubDomain: v.SubDomain,
				Value:     v.Value,
			}
			break
		}
	}

	//
	if p.dataRecord == nil {
		// this can create
		//

		return fmt.Errorf(
			"the is no this record {sub_domain: \"%s\" , line: \"%s\"}",
			p.config.SubDomain,
			line,
		)
	}

	return nil
}

func LoadFromDriverConfig(conf config.Driver) (driver.IDriver, error) {

	c := new(Client)
	c.config = conf

	c.requestClient = func() *request.HttpClient {
		clientConf := configClient.NewConfig(conf.Username, conf.Password)
		client := request.NewHttpClient(&clientConf)
		return client
	}

	c.requestClient()

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

type Domain struct {
	ID     int
	Domain string
}
type Record struct {
	ID        int
	SubDomain string
	Value     string
}

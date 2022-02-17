package record

import (
	"encoding/json"
	"fmt"
	"strings"
)

// quote: https://docs.dnspod.cn/api/5f562b21e75cf42d25bf68b6/
// 没有ipv6 ??

var _ IAdapter = new(DdnsAdapter)

type DdnsAdapter struct {
	// common.AAdapter
	adapterAbstract

	domainField
	recordLineField

	RecordID  int
	SubDomain string
	value     string

	response DdnsResponse
}

// SetValue recordType if == 6 then ipv6 else ipv4
func (p *DdnsAdapter) SetValue(value string) {
	p.value = value
}
func (p *DdnsAdapter) Method() string {
	return "Record.Ddns"
}

func (p *DdnsAdapter) Check() error {

	// domain
	if err := p.domainField.check(); err != nil {
		return err
	}
	p.Set("domain", p.domainField.domain)
	p.Set("domain_id", fmt.Sprintf("%d", p.domainField.domainID))

	// record_id
	if p.RecordID <= 0 {
		return fmt.Errorf("%s 不能为空", "RecordID")
	}
	p.Set("record_id", fmt.Sprintf("%d", p.RecordID))

	// sub_domain
	subDomain := strings.Trim(p.SubDomain, " ")
	if len(subDomain) == 0 {
		return fmt.Errorf("%s 不能为空", "SubDomain")
	}
	p.Set("sub_domain", subDomain)

	// record_line
	if err := p.recordLineField.check(); err != nil {
		return err
	}
	p.Set("record_line", p.recordLineField.recordLine)
	p.Set("record_line_id", fmt.Sprintf("%d", p.recordLineField.recordLineID))

	// value
	value := strings.Trim(p.value, " ")
	if len(value) == 0 {
		return fmt.Errorf("%s 不能为空", "Value")
	}
	p.Set("value", value)

	return nil
}

func (p *DdnsAdapter) Response() DdnsResponse {
	return p.response
}

func (p *DdnsAdapter) SetResponseJson(jsonBuf []byte) (err error) {
	res := DdnsResponse{}

	err = json.Unmarshal(jsonBuf, &res)
	if err != nil {
		err = fmt.Errorf("set response error: %v", err)
		return
	}

	p.response = res
	return
}

type DdnsResponse struct {
	Status Status `json:"status"`
	Record struct {
		ID        json.Number `json:"id"`
		SubDomain string      `json:"name"`
		Value     string      `json:"value"`
	} `json:"record"`
}

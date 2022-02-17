package record

import (
	"encoding/json"
	"fmt"
)

// quote: https://docs.dnspod.cn/api/5f562ae4e75cf42d25bf689e/

var _ IAdapter = new(ListAdapter)

type ListAdapter struct {
	// common.AAdapter
	adapterAbstract

	// Domain string
	domainField

	Offset uint

	Length uint

	//sub_domain
	SubDomain string

	response ListResponse
}

func (p *ListAdapter) Method() string {
	return "Record.List"
}

func (p *ListAdapter) Check() error {

	// dm := strings.Trim(p.Domain, " ")
	// if len(dm) == 0 {
	// 	return fmt.Errorf("%s 不能为空", "domain")
	// }
	// p.Set("domain", dm)

	if err := p.domainField.check(); err != nil {
		return err
	}
	p.Set("domain", p.domainField.domain)
	p.Set("domain_id", fmt.Sprintf("%d", p.domainField.domainID))

	if p.Offset > 0 {
		p.Set("offset", fmt.Sprintf("%d", p.Offset))
	}

	if p.Offset > 0 {
		p.Set("length", fmt.Sprintf("%d", p.Length))
	}

	if len(p.SubDomain) > 0 {
		p.Set("sub_domain", p.SubDomain)
	}
	// p.Set("record_type", p.SubDomain)
	// p.Set("record_line", p.SubDomain)
	// p.Set("record_line", p.SubDomain)
	// p.Set("record_line_id", p.SubDomain)
	// p.Set("keyword", p.SubDomain)

	return nil
}

func (p *ListAdapter) Response() ListResponse {
	return p.response
}

func (p *ListAdapter) SetResponseJson(jsonBuf []byte) (err error) {
	res := ListResponse{}

	err = json.Unmarshal(jsonBuf, &res)
	if err != nil {
		err = fmt.Errorf("set response error: %v", err)
		return
	}

	p.response = res
	return
}

type ListResponse struct {
	Status  Status   `json:"status"`
	Domain  Domain   `json:"domain"`
	Info    Info     `json:"info"`
	Records []Record `json:"records"`
}

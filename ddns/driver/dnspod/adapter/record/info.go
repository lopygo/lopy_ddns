package record

import (
	"encoding/json"
	"fmt"
	"strings"
)

// quote: https://docs.dnspod.cn/api/5f562affe75cf42d25bf68a9/

var _ IAdapter = new(DdnsAdapter)

type InfoAdapter struct {
	// common.AAdapter
	adapterAbstract

	domainField

	RecordID int

	Remark string

	response InfoResponse
}

func (p *InfoAdapter) Method() string {
	return "Record.Info"
}

func (p *InfoAdapter) Check() error {

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

	// remark
	remark := strings.Trim(p.Remark, " ")
	if len(remark) > 0 {
		p.Set("remark", remark)
	}

	return nil
}

func (p *InfoAdapter) Response() InfoResponse {
	return p.response
}

func (p *InfoAdapter) SetResponseJson(jsonBuf []byte) (err error) {
	res := InfoResponse{}

	err = json.Unmarshal(jsonBuf, &res)
	if err != nil {
		err = fmt.Errorf("set response error: %v", err)
		return
	}

	p.response = res
	return
}

type InfoResponse struct {
	Status Status `json:"status"`
	Domain struct {
		ID       json.Number `json:"id"`
		Domain   string      `json:"domain"`
		Grade    string      `json:"domain_grade"`
		DnspodNs []string    `json:"dnspod_ns"`
	} `json:"domain"`
	Record struct {
		ID        json.Number `json:"id"`
		SubDomain string      `json:"sub_domain"`
		Type      string      `json:"record_type"`
		Line      string      `json:"record_line"`
		LineID    string      `json:"record_line_id"`
		Value     string      `json:"value"`
		Weight    json.Number `json:"weight"`
		Mx        string      `json:"mx"`
		TTL       string      `json:"ttl"`
		Enabled   string      `json:"enabled"`
		Remark    string      `json:"remark"`
		UpdatedOn string      `json:"updated_on"`
		DomainID  json.Number `json:"domain_id"`
	} `json:"record"`
}

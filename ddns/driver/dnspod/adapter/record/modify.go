package record

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strings"
// )

// // quote: https://docs.dnspod.cn/api/5f562a49e75cf42d25bf6872/

// var _ IAdapter = new(ModifyAdapter)

// type ModifyAdapter struct {
// 	// common.AAdapter
// 	adapterAbstract

// 	DomainID     string
// 	RecordID     int
// 	recordType   string
// 	SubDomain    string
// 	RecordLine   string
// 	RecordLineID string
// 	value        string

// 	response ListResponse
// }

// // SetValue recordType if == 6 then ipv6 else ipv4
// func (p *ModifyAdapter) SetValue(value string, recordType int) {
// 	p.value = value
// 	if recordType == 6 {
// 		p.recordType = "AAAA"
// 	} else {
// 		p.recordType = "A"
// 	}
// }
// func (p *ModifyAdapter) Method() string {
// 	return "Record.Modify"
// }

// func (p *ModifyAdapter) Check() error {

// 	// domain_id
// 	dmId := strings.Trim(p.DomainID, " ")
// 	if len(dmId) == 0 {
// 		return fmt.Errorf("%s 不能为空", "DomainID")
// 	}
// 	p.Set("domain_id", dmId)

// 	// record_id
// 	if p.RecordID <= 0 {
// 		return fmt.Errorf("%s 不能为空", "RecordID")
// 	}
// 	p.Set("record_id", fmt.Sprintf("%d", p.RecordID))

// 	// sub_domain
// 	subDomain := strings.Trim(p.SubDomain, " ")
// 	if len(subDomain) == 0 {
// 		return fmt.Errorf("%s 不能为空", "SubDomain")
// 	}
// 	p.Set("sub_domain", subDomain)

// 	// record_type
// 	recordType := strings.Trim(p.recordType, " ")
// 	if len(recordType) == 0 {
// 		return fmt.Errorf("%s 不能为空", "RecordType")
// 	}
// 	p.Set("record_type", recordType)

// 	// record_line
// 	recordLine := strings.Trim(p.RecordLine, " ")
// 	if len(recordLine) == 0 {
// 		return fmt.Errorf("%s 不能为空", "RecordLine")
// 	}
// 	p.Set("record_line", recordLine)

// 	// record_line_id
// 	recordLineId := strings.Trim(p.RecordLineID, " ")
// 	if len(recordLineId) == 0 {
// 		return fmt.Errorf("%s 不能为空", "RecordLineID")
// 	}
// 	p.Set("record_line_id", recordLineId)

// 	// value
// 	value := strings.Trim(p.value, " ")
// 	if len(value) == 0 {
// 		return fmt.Errorf("%s 不能为空", "Value")
// 	}
// 	p.Set("value", value)

// 	// 设一个默认的
// 	p.Set("mx ", "5")

// 	return nil
// }

// func (p *ModifyAdapter) Response() ListResponse {
// 	return p.response
// }

// func (p *ModifyAdapter) SetResponseJson(jsonBuf []byte) (err error) {
// 	res := ListResponse{}

// 	err = json.Unmarshal(jsonBuf, &res)
// 	if err != nil {
// 		err = fmt.Errorf("set response error: ", err)
// 		return
// 	}

// 	p.response = res
// 	return
// }

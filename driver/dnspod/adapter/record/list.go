package record

import (
	"fmt"
	"strings"

	"github.com/lopygo/lopy_ddns/driver/dnspod/common"
)

// quote: https://docs.dnspod.cn/api/5f562ae4e75cf42d25bf689e/

var _ common.IAdapter = new(ListAdapter)

type ListAdapter struct {
	common.AAdapter

	Domain string

	Offset uint

	Length uint

	//sub_domain
	SubDomain string
}

func (p *ListAdapter) Method() string {
	return "Record.List"
}

func (p *ListAdapter) Check() error {

	dm := strings.Trim(p.Domain, " ")
	if len(dm) == 0 {
		return fmt.Errorf("%s 不能为空", "domain")
	}
	p.Set("domain", dm)

	if p.Offset > 0 {
		p.Set("offset", fmt.Sprintf("%d", p.Offset))
	}

	if p.Offset > 0 {
		p.Set("length", fmt.Sprintf("%d", p.Length))
	}

	return nil
}

// func (p *ListAdapter) GetData() map[string]string {
// 	panic("not implemented") // TODO: Implement
// }

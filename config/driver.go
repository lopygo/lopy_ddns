package config

import "strings"

const (
	IPTypeV4 int = 4
	IPTypeV6 int = 6
)

type Driver struct {
	Driver    string
	Username  string
	Password  string
	Domain    string
	SubDomain string
	IPType    string
	Ext       map[string]interface{}
}

func (p *Driver) GetType() int {
	theType := strings.ToLower(p.IPType)
	if theType == "ipv6" || theType == "v6" || theType == "6" {
		return IPTypeV6
	}

	return IPTypeV4
}

package config

import "strings"

type Driver struct {
	Driver    string
	Username  string
	Password  string
	Domain    string
	SubDomain string
	Type      string
	Ext       map[string]interface{}
}

func (p *Driver) GetType() int {
	theType := strings.ToLower(p.Type)
	if theType == "ipv6" || theType == "v6" || theType == "6" {
		return 6
	}

	return 4
}

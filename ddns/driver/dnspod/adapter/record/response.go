package record

import "encoding/json"

type Status struct {
	Code      json.Number `json:"code"`
	Message   string      `json:"message"`
	CreatedAt string      `json:"created_at"`
}

type Domain struct {
	ID       json.Number `json:"id"`
	Domain   string      `json:"name"`
	Grade    string      `json:"grade"`
	DnspodNs []string    `json:"dnspod_ns"`
}

type Info struct {
	SubDomains  string `json:"sub_domains"`
	RecordTotal string `json:"record_total"`
	RecordsNum  string `json:"records_num"`
}

type Record struct {
	ID        json.Number `json:"id"`
	Value     string      `json:"value"`
	SubDomain string      `json:"name"`
	TTL       string      `json:"ttl"`
	Enabled   string      `json:"enabled"`
	UpdatedOn string      `json:"updated_on"`
	Line      string      `json:"line"`
	LineID    string      `json:"line_id"`
	Type      string      `json:"type"`
	Weight    json.Number `json:"weight"`
	Remark    string      `json:"remark"`
	Mx        string      `json:"mx"`
}

type Response struct {
	Status Status `json:"status"`
}

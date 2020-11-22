package response

import (
	"encoding/json"
)

type DomainInfo struct {
	ID   json.Number `json:"id"`
	Name string      `json:"name"`
	// Punycode  string   `json:"punycode"`
	// Grade     string   `json:"grade"`
	// Owner     string   `json:"owner"`
	// ExtStatus string   `json:"ext_status"`
	// TTL       int      `json:"ttl"`
	// DnspodNs  []string `json:"dnspod_ns"`
}

type ListResult struct {
	Status  Status       `json:status`
	Domain  DomainInfo   `json:domain`
	Records []RecordInfo `json:records`
}

func ListResultFromBuffer(buf []byte) (ListResult, error) {
	r := ListResult{}
	err := json.Unmarshal(buf, &r)
	return r, err
}

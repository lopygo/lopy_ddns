package response

import "encoding/json"

// 暂时有用的，只有这几个
type RecordInfo struct {
	ID     json.Number `json:"id"`
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Value  string      `json:"value"`
	Type   string      `json:"type"`
	Line   string      `json:"line"`
	LineID string      `json:"line_id"`
	// TTL           string      `json:"ttl"`
	// Weight        interface{} `json:"weight"`
	// Mx            string      `json:"mx"`
	// Enabled       string      `json:"enabled"`
	// MonitorStatus string      `json:"monitor_status"`
	// Remark        string      `json:"remark"`
	// UpdatedOn     string      `json:"updated_on"`
	// UseAqb        string      `json:"use_aqb"`
}

type ModifyResult struct {
	Status Status     `json:"status"`
	Record RecordInfo `json:"record"`
}

func ModifyResultFromBuffer(buf []byte) (ModifyResult, error) {
	r := ModifyResult{}
	err := json.Unmarshal(buf, &r)
	return r, err
}

package response

import (
	"reflect"
	"testing"
)

func TestListResultFromBuffer(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ListResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "aha",
			args: args{
				// demo copy from https://docs.dnspod.cn/api/5f562ae4e75cf42d25bf689e/
				buf: []byte(`
				{
					"status": {
						"code": "1",
						"message": "Action completed successful",
						"created_at": "2018-06-11 10:41:18"
					},
					"domain": {
						"id": "12600793",
						"name": "example.com",
						"punycode": "example.com",
						"grade": "DP_Free",
						"owner": "mailbox@example.com",
						"ext_status": "dnserror",
						"ttl": 600,
						"dnspod_ns": [
							"ns3.dnsv5.com",
							"ns4.dnsv5.com"
						]
					},
					"info": {
						"sub_domains": "7",
						"record_total": "4",
						"records_num": "3"
					},
					"records": [
						{
							"id": "13608148",
							"name": "www",
							"line": "电信",
							"line_id": "10=0",
							"type": "A",
							"ttl": "600",
							"value": "1.10.0.3",
							"weight": null,
							"mx": "0",
							"enabled": "1",
							"status": "enabled",
							"monitor_status": "",
							"remark": "",
							"updated_on": "2018-06-11 10:12:51",
							"use_aqb": "no"
						}
					]
				}
				`),
			},
			want: ListResult{
				Status: Status{
					Code:      "1",
					Message:   "Action completed successful",
					CreatedAt: "2018-06-11 10:41:18",
				},
				Domain: DomainInfo{
					ID:   "12600793",
					Name: "example.com",
				},
				Records: []RecordInfo{
					RecordInfo{
						ID:     "13608148",
						Name:   "www",
						Type:   "A",
						Status: "enabled",
						Value:  "1.10.0.3",
						Line:   "电信",
						LineID: "10=0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListResultFromBuffer(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListResultFromBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListResultFromBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

package common

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsIpv4(t *testing.T) {
	Convey("default", t, func() {

		caseList := []testIsIPTypeArgs{
			testIsIPTypeArgs{
				title:  "ipv4",
				input:  "127.0.0.1",
				expect: true,
			},
			testIsIPTypeArgs{
				title:  "ipv4",
				input:  "0.0.0.0",
				expect: true,
			},
			testIsIPTypeArgs{
				title:  "ipv4",
				input:  "0.0.0.1",
				expect: true,
			},
			testIsIPTypeArgs{
				title:  "ipv4",
				input:  "255.255.255.255",
				expect: true,
			},
			testIsIPTypeArgs{
				title:  "ipv4",
				input:  "255.255.255.256",
				expect: false,
			},
		}

		for k, v := range caseList {

			Convey(fmt.Sprintf("ipv4 test %d", k+1), func() {
				So(IsIpv4(v.input), ShouldEqual, v.expect)
			})
		}

	})
}

func TestIsIpv6(t *testing.T) {
	Convey("default", t, func() {

		caseList := []testIsIPTypeArgs{
			testIsIPTypeArgs{
				title:  "ipv6",
				input:  "2408:8262:188c:544c:a401:d4ff:fef5:737f",
				expect: true,
			},
		}

		for k, v := range caseList {

			Convey(fmt.Sprintf("ipv4 test %d", k+1), func() {
				So(IsIpv6(v.input), ShouldEqual, v.expect)
			})
		}

	})
}

type testIsIPTypeArgs struct {
	title  string
	input  string
	expect bool
}

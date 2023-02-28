package members_3322_org

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIpDriver_Resolve(t *testing.T) {
	Convey("default", t, func() {

		Convey("ip length", func() {
			ipDriver := &IpDriver{}
			r, err := ipDriver.Resolve()

			So(err, ShouldBeNil)

			t.Logf("ip: %s", r)
			So(len(r), ShouldBeLessThanOrEqualTo, 15)
			So(len(r), ShouldBeGreaterThanOrEqualTo, 7)
		})

	})
}

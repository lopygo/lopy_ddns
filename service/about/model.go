package about

import (
	"fmt"
	"runtime"
)

type Info struct {
	// AppName name of your app
	AppName string

	// WebSite url of your website
	WebSite string

	// AppVersion your add AppVersion
	AppVersion string

	// BuildTime time of building your app
	BuildTime string

	// GoVersion cmd "go version" of building env
	BuildGoVersion string

	// git rev-parse HEAD
	GITCommit string

	// GoVersion cmd "go version" of building env
	GoVersion string
}

func FromInput() (p Info, err error) {
	p.AppName = appName
	p.WebSite = webSite
	p.AppVersion = appVersion
	p.BuildTime = buildTime
	p.BuildGoVersion = buildGoVersion
	p.GITCommit = gitCommit
	p.GoVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	return
}

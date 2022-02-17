export GO111MODULE=on
export CGO_ENABLED=0
export APP_VERSION ?= version-test

distDir := ./dist
pkgName := github.com/lopygo/lopy_ddns
appName := lopy_ddns
buildDate := `date +%Y-%m-%d" "%H:%M:%S`
goVersion := `go version|awk '{print $$3,$$4}'`
gitCommit := `git rev-parse HEAD`
webSite := lopygo.com

all: 

	# ready: generate builder use to generate build shell

	# build: build app to bin

	# test: unit test

	# dev: use to env of develop 

	# fmt: fmt source code 

	# clean: fmt source code 

ready: buildpre gogen

build: gobuild

test: gotest


buildpre:

	go mod tidy

gobuild:

	mkdir dist -p
	go build -ldflags " \
		-w -s \
		-X \"${pkgName}/service/about.appName=${appName}\" \
		-X \"${pkgName}/service/about.appVersion=${APP_VERSION}\" \
		-X \"${pkgName}/service/about.buildTime=${buildDate}\" \
		-X \"${pkgName}/service/about.buildGoVersion=${goVersion}\" \
		-X \"${pkgName}/service/about.gitCommit=${gitCommit}\" \
		-X \"${pkgName}/service/about.webSite=${webSite}\" \
	" \
	-o ./dist/lopy_ddns \
	./cmd/client/main.go

gogen:

	go generate ./...

gotest:

	go test -v --cover ./...
	
fmt:

	go fmt ./...
	
clean:

	rm -rf $(distDir)/*

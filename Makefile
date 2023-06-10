.PHONY: build

serviceName=remembrall
version=$(shell git describe --abbrev=0 --tags || echo '0.0.0')
versionFile=$(shell echo $(version) | tr . _)
versionFlag="main.Version=$(version)"
timeFlag="main.BuildTime=$(shell date +'%d/%m/%y_%H:%M')"


build:
	go build -ldflags="-X $(versionFlag) -X $(timeFlag)" -o bin/$(serviceName) main.go
	GOARCH=amd64 GOOS=linux go build -ldflags="-X $(versionFlag) -X $(timeFlag)" -o bin/$(serviceName)-linux-amd64-$(versionFile) main.go
	GOARCH=amd64 GOOS=darwin go build -ldflags="-X $(versionFlag) -X $(timeFlag)" -o bin/$(serviceName)-darwin-amd64-$(versionFile) main.go
	GOARCH=amd64 GOOS=windows go build -ldflags="-X $(versionFlag) -X $(timeFlag)" -o bin/$(serviceName)-windows-amd64-$(versionFile).exe main.go

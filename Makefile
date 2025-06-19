local:
	go mod tidy
	go build -v \
	-o bin/mxdemo *.go

mac-arm:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -v \
	-o bin/mxdemo-mac-arm *.go

mac-intel:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -v \
	-o bin/mxdemo-mac-intel *.go

windows-intel:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/mxdemo.exe *.go

linux-intel:
	env GOOS=linux GOARCH=amd64 go build -v -o mxdemo *.go
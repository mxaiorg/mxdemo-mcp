local:
	go mod tidy
	go build -v \
	-o bin/mxdemo *.go
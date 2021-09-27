export version=1.0
build:
	echo "start build"
	#windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/64/fserver.exe main.go
    #linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/64/fserver main.go
    #mac
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/mac/64/fserver main.go
install:
	go install
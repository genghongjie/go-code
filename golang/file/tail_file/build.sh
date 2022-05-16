#windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o _out/app-lib.exe main.go

#linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o _out/app-lib main.go

#mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o _out/app-libg-mac main.go

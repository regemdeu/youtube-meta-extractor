.PHONY: build build-windows

build:
	go build -o ./build/yke ./cmd/main.go
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/yke.exe ./cmd/main.go
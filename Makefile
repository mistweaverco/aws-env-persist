build-linux-32:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/aws-env-persist-linux-386 src/*.go
build-linux-64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/aws-env-persist-linux src/*.go
build-macos-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/aws-env-persist-mac-arm64 src/*.go

build: build-linux-32 build-linux-64 build-macos-arm64

optimize:
	upx --best --lzma dist/*

build-and-install: build-linux-64
	upx --best --lzma dist/aws-env-persist-linux
	sudo cp dist/aws-env-persist-linux /usr/local/bin/aws-env-persist

run-get-env:
	go run ./src/*.go -- get-env

run-save:
	go run ./src/*.go -- save


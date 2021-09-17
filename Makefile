build:
	go build -ldflags="-s -w"
	upx limanctl
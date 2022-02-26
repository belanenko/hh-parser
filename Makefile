run:
	clear 
	go run ./cmd/app/app.go --threads 5 --pfp=/home/tim/code/github.com/belanenko/hh-parser/assets/socks5.txt --startid=1000 --countid=50 --clickconf /home/tim/.config/hh-parser/config.json
test: 
	clear
	go test -v ./... -timeout 30s
.PHONY:build
build:
	clear 
	go build ./cmd/app/app.go

DEFAULT_GOAL=run

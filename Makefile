run: 
	clear 
	go run ./cmd/app/app.go
runf: 
	clear 
	go run ./cmd/app/app.go --threads 5 --pfp=/home/tim/code/github.com/belanenko/hh-parser/assets/socks5.txt --startid=1000 --countid=50
test: 
	clear
	go test -v ./... -timeout 30s
DEFAULT_GOAL=run

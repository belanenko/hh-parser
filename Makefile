run: 
	clear 
	go run ./cmd/app/app.go
test: 
	go test -v ./... -timeout 30s
DEFAULT_GOAL=run

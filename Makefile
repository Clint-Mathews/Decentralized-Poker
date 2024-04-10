# @ - Infront of command will not print command on make command execution
build:
	@go build -o bin/decentralized-poker

run: build
	@./bin/decentralized-poker

test: 
	go test -v ./...
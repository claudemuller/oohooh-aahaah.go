build:
	go build -o ohah cmd/interpreter/main.go

repl:
	go run cmd/repl/main.go

test:
	go test ./...

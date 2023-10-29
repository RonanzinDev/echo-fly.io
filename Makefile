build:
	go build -o ./bin/app ./cmd/main/main.go 

run: build
	./bin/app
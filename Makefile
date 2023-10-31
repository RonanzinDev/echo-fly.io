run: build
	./bin/app

build:
	go build -o ./bin/app ./cmd/main/main.go 

reload: 
	./bin/air -c air.toml
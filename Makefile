run: build
	./bin/app

build:
	go build -o ./bin/app ./main.go 

reload: 
	./bin/air -c air.toml
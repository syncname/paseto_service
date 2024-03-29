test:
	go test -v -cover ./...

build:
	go build

build-and-run: build
	./pasetoservice

build:
	@go build -C gobank -o bin/gobank


run: build
	@./gobank/bin/gobank

test:
	@go test -v ./..

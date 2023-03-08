VERSION=0.1.0

all: gen lint vet test build

build: tt

tt:
	@cd cmd/$@ && go build -o ../../bin/$@

test:
	@go test ./...

vet:
	@go vet ./...

lint:
	@revive ./...

gen:
	@buf generate

clean:
	@rm -rf pkg/tt
	@rm -rf doc/swagger

buf-lint:
	@buf lint

publish: gen test
	@buf push proto

count:
	@echo "Linecounts excluding generated and third party code"
	@gocloc --not-match-d='pkg/tt|doc/swagger' .

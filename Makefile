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
	@gocloc --not-match-d='pkg/dx|doc/swagger' .

docker-image: gen test
	@echo "Cross compiling"
	@cd cmd/dx && GOOS=linux GOARCH=amd64 go build -o ../../bin/dx-linux --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@docker build -t dx . && \
		docker tag dx:latest ghcr.io/lab5e/dx:$(VERSION)
		docker tag dx:latest ghcr.io/lab5e/dx:latest

docker-push:
	@echo "Pushing new docker image"
	@docker push ghcr.io/lab5e/dx:$(VERSION)
	@docker push ghcr.io/lab5e/dx:latest

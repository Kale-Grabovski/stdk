all: test build

test:
	go test ./...

build:
	go build src/cmd/crm/crm.go && mv crm .bin/ && go build src/cmd/api/api.go && mv api .bin/

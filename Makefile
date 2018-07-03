BINARY=abfeature
MAIN=cmd/ABFeature/main.go
VERSION=0.0.1
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
MOCK=mockery

.PHONY: clean test docs

default: build

build: test
	go build $(LDFLAGS) -o $(BINARY) $(MAIN)

build-darwin: test
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY)-darwin $(MAIN)

build-linux: test
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY)-linux $(MAIN)

build-windows: test
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY)-windows $(MAIN)

all: clean test build-darwin build-linux build-windows

test:
	go test -v ./...

run: test
	go run $(MAIN)

docker:
	docker build -t $(BINARY):$(VERSION) .

docs:
	godoc -http=:6061

fmt:
	go fmt ./...

clean:
	go clean
	rm -rf bin/$(BINARY)*
	rm -f $(BINARY)

# Project Specific
gen-mocks:
	$(MOCK) -name=IProjectStore -dir=./models/services
	$(MOCK) -name=IExperimentStore -dir=./models/services

encrypt-config:
	gpg -c config/config.yaml

gen-config:
	gpg config/config.yaml.gpg

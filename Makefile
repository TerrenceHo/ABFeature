GO=go
BINARY=abfeature
MAIN=cmd/ABFeature/main.go
VERSION=0.0.2
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
MOCK=mockery

.PHONY: clean test docs

default: install

build: test
	$(GO) build $(LDFLAGS) -o $(BINARY) $(MAIN)

build-darwin: test
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-darwin $(MAIN)

build-linux: test
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-linux $(MAIN)

build-windows: test
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-windows $(MAIN)

all: clean test build-darwin build-linux build-windows

test:
	$(GO) test -v ./...

run: test
	$(GO) run $(MAIN)

install: test
	$(GO) install $(LDFLAGS) ./...

docker:
	docker build -t $(BINARY):$(VERSION) .

docs:
	godoc -http=:6061

fmt:
	$(GO) fmt ./...

clean:
	$(GO) clean
	rm -rf bin/$(BINARY)*
	rm -f $(BINARY)

# Project Specific
gen-mocks:
	$(MOCK) -name=IProjectStore -dir=./services
	$(MOCK) -name=IExperimentStore -dir=./services

encrypt-config:
	gpg -c config/config.yaml

gen-config:
	gpg config/config.yaml.gpg

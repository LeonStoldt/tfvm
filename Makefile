PROJECT_DIR ::= ${PWD}

all:	test build lint build_windows_and_macosx

lint:
	go get -u golang.org/x/lint/golint
	golint ./... 

build:
	go vet ./...
	mkdir -p target/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags \"-static\"' -o target/tfvm-linux_amd64 cmd/main.go
	cp target/tfvm-linux_amd64 tfvm

format:
	go fmt ./...

build_windows_and_macosx:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-extldflags \"-static\"' -o target/tfvm-windows_amd64 cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -ldflags '-extldflags \"-static\"' -o target/tfvm-windows_386 cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags \"-static\"' -o target/tfvm-darwin_amd64 cmd/main.go

test:
	go test ./...

build_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:1.13-stretch \
		make build

test_with_docker:
	docker run -u $(shell id -u):$(shell id -g) \
		-v ${PROJECT_DIR}:/build \
		-w /build \
		-e HOME=/build \
		golang:1.13-stretch \
		make test

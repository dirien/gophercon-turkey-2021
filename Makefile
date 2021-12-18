VERSION=0.0.1
BUILD=`git rev-parse --short HEAD`
BINARY_NAME=gophercon-turkey-2021
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.build=${BUILD}"
DIST=dist

.PHONY: all
all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "build          - runs go build"
	@echo "build_all      - runs go build with ldflags version=${VERSION} & build=${BUILD}"
	@echo "docker         - builds docker image"
	@echo ""

build: clean
	@echo "building..."
	@go build ${LDFLAGS} -o ${DIST}/${BINARY_NAME} .

build_all: clean
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build ${LDFLAGS} -v -o ${DIST}/$(BINARY_NAME)-$(GOOS)-$(GOARCH))))

clean:
	@echo "cleaning..."
	@go clean
	@rm -rf ${DIST}

test:
	@echo "testing..."
	@go test -v ./...

lint:
	@echo "linting..."
	@golangci-lint run ./...

docker: clean build_all
	@echo "building docker image..."
	@docker build -t ${BINARY_NAME}:${VERSION} .
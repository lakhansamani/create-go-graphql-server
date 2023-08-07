DEFAULT_VERSION=0.1.0-local
VERSION := $(or $(VERSION),$(DEFAULT_VERSION))
DOCKERIMAGE=app # change this to your docker image name

all: binaries
run:
	go run -ldflags "-w -X main.version=$(VERSION)" main.go
clean:
	rm -rf bin
binaries:
	gox \
		-osarch="linux/amd64 linux/arm64 darwin/amd64 windows/amd64" \
		-ldflags "-w -X main.version=$(VERSION)" \
		-output="bin/{{.OS}}/{{.Arch}}/app" \
		./...
docker: binaries
	docker context create multiplatform
	docker buildx install
	docker buildx create --name multiplatform --use multiplatform
	docker build --platform linux/amd64,linux/arm64 --push -t $(DOCKERIMAGE) .
docker-local: binaries
	docker build -t $(DOCKERIMAGE) .
generate-graphql:
	go get github.com/99designs/gqlgen@v0.17.36 && go run github.com/99designs/gqlgen generate
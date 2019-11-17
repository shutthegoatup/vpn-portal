VERSION := "0.2"
GIT_COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: all test build

all: build

build:
	go build -o portal ./cmd/vpn-portal/...

lint:
	go fmt 

container:
	/kaniko/executor 
		--context $CI_PROJECT_DIR 
		--dockerfile $CI_PROJECT_DIR/build/package/Dockerfile
		--destination $CI_REGISTRY_IMAGE:$CI_COMMIT_TAG
		--cache=false
		--reproducible=true
		--digest-file=digest.txt

clean:
	rm vpn-portal


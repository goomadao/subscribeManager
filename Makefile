export PATH := $(PATH):$(GOPATH)/bin
.PHONY: all build test client server web clean

all: build

build: 
	go get ./...
	go build -o bin/subscribeManager ./main.go

web:
	cd web; \
	yarn; \
	yarn build; \
	cd ..; \
	statik -src=web/dist

clean:
	rm -rf bin
	rm -rf web/dist
	rm -rf web/node_modules
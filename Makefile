.PHONY: dep build fmt

dep:
		dep ensure

build: dep
		go build -o todos .

fmt: dep
		go fmt ./...

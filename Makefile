default: y.go generate
.PHONY: default

y.go: asn1.y
	goyacc asn1.y

generate:
	go generate -v ./...
.PHONY: generate

deps:
	go get golang.org/x/tools/cmd/goyacc
.PHONY: deps

codegen: y.go
.PHONY: codegen

clean:
	rm -f y.go
	find . -name '*_generated.go' -exec rm '{}' \;
.PHONY: clean

test: default
	go test -v ./...
.PHONY: test
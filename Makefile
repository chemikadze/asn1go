default: y.go
.PHONY: default

y.go: asn1.y
	goyacc asn1.y

deps:
	go get golang.org/x/tools/cmd/goyacc
.PHONY: deps

codegen: y.go
.PHONY: codegen

clean:
	rm -f y.go
.PHONY: clean

test: y.go
	go test -v ./...
.PHONY: test
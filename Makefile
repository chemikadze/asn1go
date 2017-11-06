default: y.go
.PHONY: default

y.go: asn1.y
	goyacc asn1.y

runmain: y.go
	go run main/main.go
.PHONY: runmain

clean:
.PHONY: clean

test: y.go
	go test -v ./...
.PHONY: test
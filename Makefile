.PHONY: go
build:
	make go
go:
	./build.sh go base.proto go/base
	./build.sh go postar.proto go/postar
base:
	./build.sh go base.proto go/base
postar:
	./build.sh go postar.proto go/postar
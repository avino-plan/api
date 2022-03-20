.PHONY: go
GO_OUT=go-out
build:
	make go
go:
	./build.sh go base.proto $(GO_OUT)/base
	./build.sh go postar.proto $(GO_OUT)/postar
base:
	./build.sh go base.proto $(GO_OUT)/base
postar:
	./build.sh go postar.proto $(GO_OUT)/postar
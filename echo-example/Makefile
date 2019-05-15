.PHONY: proto
proto:
	protoc \
		-I/usr/local/include \
		-I. \
		-I$(GOPATH)/src \
		--go_out=plugins=grpc:. \
		service.proto
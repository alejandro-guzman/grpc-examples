.PHONY: proto-gw
proto:
	protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		service.proto

.PHONY: proto-rp
proto-rp:
	protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		service.proto

.PHONY: proto-sg
proto-sg:
	protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		service.proto
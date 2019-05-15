# grpc-examples

- [grpc-examples](#grpc-examples)
  - [Install depnedencies](#install-depnedencies)
    - [gRPC](#grpc)
    - [protoc](#protoc)
    - [Install protoc plugin for GO](#install-protoc-plugin-for-go)
    - [Install this package](#install-this-package)
    - [Install grpcurl](#install-grpcurl)
  - [Run the echo example](#run-the-echo-example)
  - [Run the echo gateway example](#run-the-echo-gateway-example)
    - [Install additional dependencies to run this example](#install-additional-dependencies-to-run-this-example)
      - [grpc-gateway](#grpc-gateway)
      - [protoc swagger plugin](#protoc-swagger-plugin)

## Install depnedencies

### gRPC

Use `go get` to install grpc

```bash
go get -u google.golang.org/grpc
```

### protoc

Follow this link and download protoc compiler for your arch

[protobuf releases](https://github.com/google/protobuf/releases)

Unzip the file and add it to your path

### Install protoc plugin for GO

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Install this package

```bash
go get -u github.com/agguzman/grpc-examples/echo-example
```

### Install grpcurl

```bash
brew install grpcurl
```

## Run the echo example

Open a terminal and `cd` into `grpc-examples/echo-example`

Run the server

```bash
go run server/main.go
```

Open a separate terminal and run the client

```bash
go run client/main.go "my message"
```

## Run the echo gateway example

The grpc gateway plugin uses the service definintions to create a reverse proxy and maps REST API requests to gRPC. Can optionally emit API definitions Open API v2.

### Install additional dependencies to run this example

#### grpc-gateway

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

#### protoc swagger plugin

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```


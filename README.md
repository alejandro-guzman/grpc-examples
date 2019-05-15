# grpc-examples

- [grpc-examples](#grpc-examples)
  - [Install depnedencies](#install-depnedencies)
    - [gRPC](#grpc)
    - [protoc](#protoc)
    - [Install protoc plugin for GO](#install-protoc-plugin-for-go)
    - [Install this package](#install-this-package)
    - [Install grpcurl](#install-grpcurl)
  - [Run the example](#run-the-example)

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

## Run the example

Open a terminal and `cd` into `grpc-examples/echo-example`

Run the server

```bash
go run server/main.go
```

Open a separate terminal and run the client

```bash
go run client/main.go "my message"
```
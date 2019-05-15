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
  - [Docs example](#docs-example)

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

The go files with already be generated but the command used to is

```bash
cd echo-example
make proto # generates service.pb.go
```

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

The go files with already be generated but the command used to is

```bash
cd echo-gateway-example
make proto # creates service.pb.go
make proto-rp # creates service.pb.gw.go
make proto-sg # creates service.swagger.json
```

Run the server

```bash
cd echo-gateway-examples
go run server/main.go
```

Run the reverse proxy

```bash
go run gateway/main.go
```

Hit the gateway with the client from echo-example

```bash
cd echo-example
go run client/main.go "jello"
```

## Docs example

Install the docs plugin for protoc or do what I did and just docker to run the plugin.

Take a look at `echo-docs-example/Makefile` and notice I'm usingh docker to create the docs.

```bash
cd echo-docs-example
rm -r docs/*
make proto-docs
```

You can the format of the output by include the options

```bash
--doc_opt=json,docs.json
# or
--doc_opt=markdown,docs.md
# or
--doc_opt=html,index.html
```

[Further reading](https://github.com/pseudomuto/protoc-gen-doc/)
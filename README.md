gRPC Client and Server Basic Example
====================================

There are to applications in this repo:

1. Echo-Client
2. Echo-Server

The first one sends a text line to server and expects it to return capitalized text with parentheses around. 
The second one receives a text line and responses with capitalized text with parentheses around.

How to build
------------

To run this example **from scratch** you need to install Protocol Buffer compiler. See [instructions](https://grpc.io/docs/protoc-installation/).

Delete generated files from `api` directory. 

    rm -f api/*.go

Then you need to compile .proto file.

     protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/echo.proto 

Or you could just run server and client 👻

How to run
----------

Two run server

    go run server/main.go

And then you need a separate terminal session for client

    go run client/main.go

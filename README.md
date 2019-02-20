# Simple gRPC server with REST API gateway written in Go.
A simple client server project with gRPC communication and REST API client. 

# Download and build the project  

## Downloading 
Download the project and the project dependencies

`go get github.com/romangurevitch/grpc-rest-go...`

## Testing the project:
Run the following command:

`go test -v github.com/romangurevitch/grpc-rest-go/...`
## Running the project:
In the root project path:

Running the server: 

`go run server/server.go`

Running the client:

`go run client/client.go`

## Documentation 
REST API Swagger documentation available at:

`http://localhost:8081/swagger`

# Developers guide
This project uses [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to generate both REST API and swagger configuration from protoc spec. 

All generated files are under the `api` package:

`api.pb.go`: services and messages implementation. 

To generate, under api package path:
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api.proto
```

`api.pb.gw.go` gRPC geteway implementation.

To generate, under api package path:
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
	api.proto
```

`api.swagger.json` REST API swagger definitions.

To generate, under api package path:
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  api.proto	
```




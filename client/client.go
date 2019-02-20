package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/romangurevitch/grpc-rest-go"
	"github.com/romangurevitch/grpc-rest-go/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"path"
	"strings"

	"net/http"
)

// Dictionary client.
// REST api documentation available using swagger-ui at: http://localhost:<config.HttpPort>/swagger
func main() {
	log.Println("Starting dictionary client")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// REST to gRPC handling
	err := api.RegisterDictionaryServiceHandlerFromEndpoint(ctx, gw, "localhost"+config.GrpcPort, opts)
	if err != nil {
		log.Fatal(err)
	}

	// adding swagger-ui handling
	mux := http.NewServeMux()
	mux.Handle("/api/", gw)
	mux.HandleFunc("/swagger/", serveSwagger)

	log.Println("Starting HTTP client on port " + config.HttpPort)
	log.Println("Swagger REST API documentation available at: " + "localhost" + config.HttpPort + "/swagger")
	if err = http.ListenAndServe(config.HttpPort, mux); err != nil {
		log.Fatal(err)
	}
}

func serveSwagger(responseWriter http.ResponseWriter, request *http.Request) {
	swaggerPath := strings.TrimPrefix(request.URL.Path, "/swagger/")
	swaggerPath = path.Join("swagger-ui/", swaggerPath)
	http.ServeFile(responseWriter, request, swaggerPath)
}

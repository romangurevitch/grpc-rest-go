package main

import (
	"context"
	"fmt"
	"github.com/romangurevitch/grpc-rest-go"
	"github.com/romangurevitch/grpc-rest-go/api"
	"github.com/romangurevitch/grpc-rest-go/dictionary"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Starting dictionary service")
	server, err := createDictionaryServer()
	if err != nil {
		log.Fatal(err)
	}
	defer server.GracefulStop()

	listen, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting gRPC server on port " + config.GrpcPort)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func createDictionaryServer() (*grpc.Server, error) {
	server := grpc.NewServer()
	dic, err := dictionary.New()
	if err != nil {
		return nil, err
	}
	api.RegisterDictionaryServiceServer(server, &dictionaryService{dictionary: dic})

	return server, nil
}

type dictionaryService struct {
	dictionary dictionary.Dictionary
}

func (ds *dictionaryService) FindWord(ctx context.Context, word *api.Word) (*api.FindWordRes, error) {
	logMethod("FindWord", word)
	exist, err := ds.dictionary.Find(word.Word)
	return &api.FindWordRes{Exist: exist, Result: getResult(err)}, nil
}

func (ds *dictionaryService) AddWord(ctx context.Context, word *api.Word) (*api.Result, error) {
	logMethod("AddWord", word)
	err := ds.dictionary.Add(word.Word)
	return getResult(err), nil
}

func (ds *dictionaryService) DeleteWord(ctx context.Context, word *api.Word) (*api.Result, error) {
	logMethod("DeleteWord", word)
	err := ds.dictionary.Delete(word.Word)
	return getResult(err), nil
}

func (ds *dictionaryService) GetMostPopularWords(ctx context.Context, empty *api.Empty) (*api.GetMostPopularWordsRes, error) {
	logMethod("GetMostPopularWords", empty)
	popular, err := ds.dictionary.MostPopular()
	return &api.GetMostPopularWordsRes{Words: popular, Result: getResult(err)}, nil
}

func logMethod(methodName string, params ...interface{}) {
	logMessage := fmt.Sprintf("INFO: method: %s, params: ", methodName)
	for _, v := range params {
		logMessage += fmt.Sprintf("%v ", v)
	}
	log.Printf("%s\n", logMessage)
}

func getResult(err error) *api.Result {
	if err == nil {
		return &api.Result{Status: api.Result_Success}
	}
	log.Println("ERROR:", err)
	return &api.Result{Status: api.Result_Fail, Error: err.Error()}
}

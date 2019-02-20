package main

import (
	"context"
	"github.com/romangurevitch/grpc-rest-go"
	"github.com/romangurevitch/grpc-rest-go/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// start gRPC server using main
	server, err := createDictionaryServer()
	if err != nil {
		log.Fatal(err)
	}
	defer server.GracefulStop()

	// start the dictionary server
	go func() {
		listen, err := net.Listen("tcp", config.GrpcPort)
		if err != nil {
			log.Fatal(err)
		}

		if err := server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
	waitUntilServiceIsUp()

	code := m.Run()
	os.Exit(code)
}

func TestFindWord(t *testing.T) {
	client, err := getDocumentGrpcClient()
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.FindWord(context.Background(), &api.Word{Word: "yes"})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Exist {
		t.Errorf("could not find word: 'yes' in the dictionary")
	}
}

func TestAddWord(t *testing.T) {
	client, err := getDocumentGrpcClient()
	if err != nil {
		t.Fatal(err)
	}

	addRes, err := client.AddWord(context.Background(), &api.Word{Word: "new-word"})
	if err != nil {
		t.Fatal(err)
	}

	if addRes.Status == api.Result_Fail {
		t.Fatal("failed to add new word: 'new-word' to the dictionary")
	}

	findRes, err := client.FindWord(context.Background(), &api.Word{Word: "new-word"})
	if err != nil {
		t.Fatal(err)
	}
	if !findRes.Exist {
		t.Errorf("could not find word: 'new-word' in the dictionary")
	}
}

func TestDeleteWord(t *testing.T) {
	client, err := getDocumentGrpcClient()
	if err != nil {
		t.Fatal(err)
	}

	deleteRes, err := client.DeleteWord(context.Background(), &api.Word{Word: "yes"})
	if err != nil {
		t.Fatal(err)
	}

	if deleteRes.Status == api.Result_Fail {
		t.Fatal("failed to delete word: 'yes'")
	}

	findRes, err := client.FindWord(context.Background(), &api.Word{Word: "yes"})
	if err != nil {
		t.Fatal(err)
	}
	if findRes.Exist {
		t.Errorf("found the word: 'yes' in the dictionary")
	}
}

func TestGetPopularWords(t *testing.T) {
	client, err := getDocumentGrpcClient()
	if err != nil {
		t.Fatal(err)
	}

	popularRes, err := client.GetMostPopularWords(context.Background(), &api.Empty{})
	if err != nil {
		t.Fatal(err)
	}

	if popularRes.Result.Status == api.Result_Fail {
		t.Fatal("failed to find popular words")
	}
}

func getDocumentGrpcClient() (api.DictionaryServiceClient, error) {
	conn, err := grpc.Dial("localhost"+config.GrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return api.NewDictionaryServiceClient(conn), nil
}

// naive health check impl
func waitUntilServiceIsUp() {
	client, err := getDocumentGrpcClient()
	if err != nil {
		log.Fatal(err)
	}
	for true {
		// adding existing word, waiting for no errors
		_, err = client.AddWord(context.Background(), &api.Word{Word: "yes"})
		if err == nil {
			break
		}
	}
}

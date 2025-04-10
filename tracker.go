package tracker

import (
	"context"
	"fmt"
	"log"

	"github.com/dexidp/dex/api/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewTrackerClient создает нового клиента для отслеживания событий
func NewTrackerClient(hostAndPort, caPath string) (api.DexClient, error) {
	creds, err := credentials.NewClientTLSFromFile(caPath, "")
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки сертификата: %v", err)
	}

	conn, err := grpc.Dial(hostAndPort, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %v", err)
	}
	return api.NewDexClient(conn), nil
}

// TrackEvents отслеживает события через gRPC
func TrackEvents(client api.DexClient) {
	req := &api.CreateClientReq{
		Client: &api.Client{
			Id:           "example-app",
			Name:         "Example App",
			Secret:       "ZXhhbXBsZS1hcHAtc2VjcmV0",
			RedirectUris: []string{"http://127.0.0.1:5555/callback"},
		},
	}

	resp, err := client.CreateClient(context.TODO(), req)
	if err != nil {
		log.Fatalf("Ошибка при отслеживании событий: %v", err)
	}
	fmt.Printf("Отслеженное событие: %+v\n", resp)
}

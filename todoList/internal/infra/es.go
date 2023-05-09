package infra

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var (
	ESClient *elasticsearch.Client
)

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "username",
		Password: "password",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	ESClient = es
}

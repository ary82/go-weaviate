// Set these environment variables
// WEAVIATE_URL      your Weaviate instance URL, without https prefix
// WEAVIATE_API_KEY  your Weaviate instance API key
// OPENAI_API_KEY    your OpenAI API key

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

// Create the client
func CreateClient() {
	cfg := weaviate.Config{
		Host:   os.Getenv("WEAVIATE_URL"),
		Scheme: "http",
		Headers: map[string]string{
			"X-OpenAI-Api-Key": os.Getenv("OPENAI_API_KEY"),
		},
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	live, err := client.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", live)
}

func main() {
	CreateClient()
}

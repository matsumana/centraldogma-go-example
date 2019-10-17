package dogma

import (
	"context"
	"encoding/json"
	"fmt"
)

type configFromCentralDogma struct {
	Greeting string `json:"greeting"`
}

// Watch Central Dogma file
func Watch(callback func(string)) error {
	dogmaFile := &CentralDogmaFile{
		Token:      "anonymous", // In Central Dogma Docker image, authentication is disabled so need to use anonymous token.
		BaseURL:    "http://localhost:36462",
		Project:    "fukuoka-go",
		Repo:       "demo",
		Path:       "/config.json",
		TimeoutSec: 30,
	}

	err := dogmaFile.Watch(context.Background(), func(bytes []byte) {
		config := new(configFromCentralDogma)
		if err := json.Unmarshal(bytes, config); err != nil {
			panic(err)
		}

		greeting := config.Greeting
		fmt.Printf("Fetched from Central Dogma: %s\n", greeting)

		callback(greeting)
	})
	if err != nil {
		return err
	}

	return nil
}

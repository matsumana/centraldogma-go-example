package dogma

import (
	"context"
	"encoding/json"
	"fmt"
)

// var dogmaFile *CentralDogmaFile

type configFromCentralDogma struct {
	Greeting string `json:"greeting"`
}

// // NewClient new client
// func NewClient() {
// 	dogmaFile = &CentralDogmaFile{
// 		Token:      "anonymous", // In Central Dogma Docker image, authentication is disabled so need to use anonymous token.
// 		BaseURL:    "http://localhost:36462",
// 		Project:    "fukuoka-go",
// 		Repo:       "demo",
// 		Path:       "/config.json",
// 		TimeoutSec: 30,
// 	}
// }

// // Fetch fetch from Central Dogma
// func Fetch() (string, error) {
// 	data, err := dogmaFile.Fetch(context.Background())
// 	if err != nil {
// 		return "", err
// 	}

// 	jsonBytes := ([]byte)(data)
// 	config := new(configFromCentralDogma)

// 	if err := json.Unmarshal(jsonBytes, config); err != nil {
// 		return "", err
// 	}
// 	fmt.Printf("Fetched from Central Dogma: %s\n", config.Greeting)

// 	return config.Greeting, nil
// }

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

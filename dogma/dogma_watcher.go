package dogma

import (
	"context"
	"encoding/json"
	"fmt"
)

type fetchedData struct {
	Greeting string `json:"greeting"`
}

// WatchFile watch file
func WatchFile(path string, callback func(string)) error {
	dogmaFile := &CentralDogmaFile{
		Token:      "anonymous", // In Central Dogma Docker image, need to use 'anonymous' because authentication is disabled
		BaseURL:    "http://localhost:36462",
		Project:    "fukuoka-go",
		Repo:       "demo",
		Path:       path,
		TimeoutSec: 30,
	}

	err := dogmaFile.Watch(context.Background(), func(bytes []byte) {
		data := new(fetchedData)
		if err := json.Unmarshal(bytes, data); err != nil {
			panic(err)
		}

		greeting := data.Greeting
		fmt.Printf("Fetched from Central Dogma: %s\n", greeting)

		callback(greeting)
	})
	if err != nil {
		return err
	}

	return nil
}

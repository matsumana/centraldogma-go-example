package dogma

import (
	"context"
	"sync/atomic"
	"time"

	"go.linecorp.com/centraldogma"
)

// This file is copied from https://github.com/line/centraldogma-go#example

// CentralDogmaFile represents a file in application repository, stored on Central Dogma server.
type CentralDogmaFile struct {
	client            atomic.Value
	BaseURL           string       `yaml:"base_url" json:"base_url"`
	Token             string       `yaml:"token" json:"token"`
	Project           string       `yaml:"project" json:"project"`
	Repo              string       `yaml:"repo" json:"repo"`
	Path              string       `yaml:"path" json:"path"`
	LastKnownRevision atomic.Value `yaml:"-" json:"-"`
	TimeoutSec        int          `yaml:"timeout_sec" json:"timeout_sec"`
}

func (c *CentralDogmaFile) getClientOrSet() (*centraldogma.Client, error) {
	if v, stored := c.client.Load().(*centraldogma.Client); stored {
		return v, nil
	}

	// create a new client
	dogmaClient, err := centraldogma.NewClientWithToken(c.BaseURL, c.Token, nil)
	if err != nil {
		return nil, err
	}

	// store
	c.client.Store(dogmaClient)

	return dogmaClient, nil
}

// Fetch file content from Central Dogma.
func (c *CentralDogmaFile) Fetch(ctx context.Context) (b []byte, err error) {
	dogmaClient, err := c.getClientOrSet()
	if err != nil {
		return
	}

	entry, _, err := dogmaClient.GetFile(ctx, c.Project, c.Repo, "", &centraldogma.Query{
		Path: c.Path,
		Type: centraldogma.Identity,
	})
	if err != nil {
		return
	}

	// set last known revision
	c.LastKnownRevision.Store(entry.Revision)

	b = entry.Content
	return
}

// Watch changes on remote file.
func (c *CentralDogmaFile) Watch(ctx context.Context, callback func([]byte)) error {
	dogmaClient, err := c.getClientOrSet()
	if err != nil {
		return err
	}

	ch, closer, err := dogmaClient.WatchFile(ctx, c.Project, c.Repo, &centraldogma.Query{
		Path: c.Path,
		Type: centraldogma.Identity,
	}, time.Duration(c.TimeoutSec)*time.Second)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				closer()
				return

			case changes := <-ch:
				if changes.Err == nil {
					callback(changes.Entry.Content)
				}
			}
		}
	}()

	return nil
}

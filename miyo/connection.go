package miyo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	ssdp "github.com/koron/go-ssdp"
)

type Conn struct {
	host   string
	apiKey string
}

type ConnectOption func(c *Conn)

func WithAPIKey(apiKey string) ConnectOption {
	return func(c *Conn) {
		c.apiKey = apiKey
	}
}

// Connect returns an object representing a MIYO Cube.
// If host is "", it uses FindCube() to discover the MIYO Cube with UPnP.
// If APIKey is "", it uses APIKey() to request a new API key.
func Connect(ctx context.Context, host, apiKey string) (*Conn, error) {
	if host == "" {
		var err error
		host, err = FindCube(ctx)
		if err != nil {
			return nil, fmt.Errorf("FindCube: %w", err)
		}
		log.Printf("MIYO Cube address: %q", host)
	}

	if apiKey == "" {
		var err error
		apiKey, err = APIKey(ctx, host)
		if err != nil {
			return nil, fmt.Errorf("APIKey: %w", err)
		}
		log.Printf("New MIYO API key: %q", apiKey)
	}

	return &Conn{
		host:   host,
		apiKey: apiKey,
	}, nil
}

// FindCube uses UPnP to discover a MIYO Cube on the local network and returns its address or hostname.
func FindCube(_ context.Context) (string, error) {
	const (
		timeoutSec = 3
		localAddr  = ""
	)
	services, err := ssdp.Search(ssdp.RootDevice, timeoutSec, localAddr)
	if err != nil {
		return "", fmt.Errorf("ssdp.Search: %w", err)
	}

	for _, srv := range services {
		if !strings.Contains(srv.Header().Get("Server"), "miyocube") {
			continue
		}

		u, err := url.Parse(srv.Location)
		if err != nil {
			return "", fmt.Errorf("url.Parse: %w", err)
		}

		return u.Host, nil
	}

	return "", errors.New("MIYO Cube not found")
}

package miyo

import (
	"context"
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

func Connect(ctx context.Context, host string, opts ...ConnectOption) (*Conn, error) {
	c := &Conn{
		host: host,
	}

	for _, opt := range opts {
		opt(c)
	}

	if err := c.link(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

package miyo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type linkResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error"`
	APIKey string `json:"apiKey"`
}

// APIKey requests a new API key from the MIYO cube.
// The physical button on the MIYO cube needs to be pressed before calling this function.
// It is typically used as part of a one-time setup.
func APIKey(ctx context.Context, addr string) (string, error) {
	url := "http://" + addr + "/api/link"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var lr linkResponse
	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return "", err
	}

	if lr.Status != "success" {
		return "", fmt.Errorf("/api/link: %v", lr.Error)
	}

	return lr.APIKey, nil
}

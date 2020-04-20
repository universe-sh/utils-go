package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	cdn = "cdn.universe.sh"
)

// Client struct
type Client struct {
	Context context.Context
	HTTP    *http.Client
}

func (c *Client) getBytes(address string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.HTTP.Do(req.WithContext(c.Context))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(endpoint string, v interface{}) error {
	data, err := c.getBytes(endpoint)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// New Client endpoint
func NewClient() *Client {
	return &Client{
		Context: context.TODO(),
		HTTP: &http.Client{
			Timeout: (3 * time.Second),
		},
	}
}

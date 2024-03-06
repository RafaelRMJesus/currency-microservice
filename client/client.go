package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RafaelRMJesus/currency-microservice/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	priceRes := new(types.PriceResponse)
	if err := json.NewDecoder(res.Body).Decode(priceRes); err != nil {
		return nil, err
	}
	return priceRes, nil
}

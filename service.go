package main

import (
	"context"
	"fmt"
)

// PriceFetcher é uma interface que da fetch em um price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implementa a PriceFetcher interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC":  67_947.90,
	"ETH":  3_817.57,
	"RAFA": 100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("The given Ticker (%s) is not supported", ticker)
	}

	return price, nil
}

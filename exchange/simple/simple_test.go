package simple

import (
	"fmt"
	"testing"

	"exchange/exchange"
)

func TestSimpleSelector(t *testing.T) {
	client := newClient()
	client.addRate("low", 0.5)
	client.addRate("high", 0.75)

	simpleSelector := SequentialSelector{client: client.getRate}
	rate, broker := simpleSelector.Select([]exchange.Broker{
		{"Low", "low"},
		{"High", "high"},
	}, "GBP", "EUR")

	if rate != 0.75 {
		t.Errorf("expected rate of 0.75, got %f", rate)
	}

	if broker.Name != "High" {
		t.Errorf("expected broker name of High, got %s", broker.Name)
	}
}

type client struct {
	urlsToRates map[string]float64
}

func newClient() *client {
	return &client{urlsToRates: make(map[string]float64)}
}

func (c *client) addRate(url string, rate float64) {
	c.urlsToRates[url] = rate
}

func (c *client) getRate(url, sell, buy string) (float64, error) {
	if rate, ok := c.urlsToRates[url]; ok {
		return rate, nil
	}

	return 0, fmt.Errorf("no rate defined for URL %s", url)
}
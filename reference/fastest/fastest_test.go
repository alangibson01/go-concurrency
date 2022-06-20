package fastest

import (
    "testing"
    "time"

    "exchange/exchange"
)

func TestFastestSelector(t *testing.T) {
    client := func(url, sell, buy string) (float64, error) {
        if url == "fast" {
            return 0.5, nil
        }

        time.Sleep(1 * time.Second)
        return 0.75, nil
    }

    fastestSelector := SelectFastest{client: client}
    rate, exchange := fastestSelector.Select([]exchange.Broker{
        {"Fast", "fast"},
        {"Slow", "slow"},
    }, "GBP", "EUR")

    if rate != 0.5 {
        t.Errorf("expected rate of 0.5, got %f", rate)
    }

    if exchange.Name != "Fast" {
        t.Errorf("expected exchange name of Fast, got %s", exchange.Name)
    }
}


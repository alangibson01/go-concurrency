package best

import (
    "testing"
    "time"

    "exchange/exchange"
)

func TestBestSelector(t *testing.T) {
    client := func(url, sell, buy string) (float64, error) {
        if url == "fast" {
            return 0.5, nil
        }

        time.Sleep(1 * time.Second)
        return 0.75, nil
    }

    bestSelector := SelectBest{client: client}
    rate, broker := bestSelector.Select([]exchange.Broker{
        {"Fast", "fast"},
        {"Slow", "slow"},
    }, "GBP", "EUR")

    if rate != 0.75 {
        t.Errorf("expected rate of 0.75, got %f", rate)
    }

    if broker.Name != "Slow" {
        t.Errorf("expected exchange name of Slow, got %s", broker.Name)
    }
}


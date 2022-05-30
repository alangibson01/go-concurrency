package simple

import "exchange/exchange"

type SequentialSelector struct {
    client exchange.BrokerAPI
}

func New() exchange.BrokerSelector {
    return &SequentialSelector{client: exchange.GetPrice}
}

func (s *SequentialSelector) Select(exchanges []exchange.Broker, sell, buy string) (float64, exchange.Broker) {
    var bestBroker exchange.Broker
    var bestRate float64
    for _, e := range exchanges {
        r, err := s.client(e.URL, sell, buy)
        if err != nil {
            continue
        }

        if r > bestRate {
            bestRate = r
            bestBroker = e
        }
    }

    return bestRate, bestBroker
}

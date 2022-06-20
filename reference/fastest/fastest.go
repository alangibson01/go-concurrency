package fastest

import "exchange/exchange"

type SelectFastest struct {
    client exchange.BrokerAPI
}

type result struct {
    rate     float64
    broker exchange.Broker
}

func (s *SelectFastest) Select(brokers []exchange.Broker, sell, buy string) (float64, exchange.Broker) {
    resultChan := make(chan result)

    for _, b := range brokers {
        go s.getPrice(b, sell, buy, resultChan)
    }

    result := <-resultChan
    return result.rate, result.broker
}

func (s *SelectFastest) getPrice(broker exchange.Broker, sell, buy string, resultChan chan result) {
    rate, err := s.client(broker.URL, sell, buy)
    if err != nil {
        return
    }

    resultChan <- result{rate: rate, broker: broker}
}

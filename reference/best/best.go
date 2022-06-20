package best

import (
    "sync"

    "exchange/exchange"
)

type SelectBest struct {
    client exchange.BrokerAPI
    result result
}

type result struct {
    rate     float64
    exchange exchange.Broker
    sync.Mutex
}

func (s *SelectBest) Select(brokers []exchange.Broker, sell, buy string) (float64, exchange.Broker) {
    var wg sync.WaitGroup

    for _, b := range brokers {
        wg.Add(1)
        go s.getPrice(b, sell, buy, &wg)
    }

    wg.Wait()
    return s.result.rate, s.result.exchange
}

func (s *SelectBest) getPrice(b exchange.Broker, sell, buy string, wg *sync.WaitGroup) {
    defer wg.Done()
    rate, err := s.client(b.URL, sell, buy)
    if err != nil {
        return
    }

    s.result.Lock()
    defer s.result.Unlock()
    if rate > s.result.rate {
        s.result.rate = rate
        s.result.exchange = b
    }
}

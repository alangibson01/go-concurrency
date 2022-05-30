package best

import "exchange/exchange"

type SelectBest struct {
    api exchange.BrokerAPI
}

func (b *SelectBest) Select(brokers []exchange.Broker, sell, buy string) (float64, exchange.Broker) {
    panic("implement me")
}

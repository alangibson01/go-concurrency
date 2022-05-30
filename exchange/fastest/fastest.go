package fastest

import "exchange/exchange"

type SelectFastest struct {
    api exchange.BrokerAPI
}

func (b *SelectFastest) Select(brokers []exchange.Broker, sell, buy string) (float64, exchange.Broker) {
    panic("implement me")
}

package main

import (
	"fmt"

	"exchange/exchange"
	"exchange/exchange/simple"
)

var brokers = []exchange.Broker{
	{"Whizz FX", "http://localhost:8080"},
	{"Steady FX", "http://localhost:9090"},
}

func main() {
	sell := "GBP"
	buy := "EUR"

	selector := simple.New()
	rate, broker := selector.Select(brokers, sell, buy)

	fmt.Printf("Rate selected is %f from %s", rate, broker.Name)
}

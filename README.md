# Concurrency in Go

Go provides a number of facilities related to concurrency, on which comprehensive documentation can be found [here](https://github.com/golang/go/wiki/LearnConcurrency).
At the very least, ["Effective Go: Concurrency"](https://go.dev/doc/effective_go#concurrency) is a must-read introduction to the topic.

This repo doesn't restate any of this information, and is instead an exercise to allow you to try out using Go's
concurrency features in a simplistic "real world" use case.

## Use Case

Your company frequently needs to make foreign exchange transactions, and has agreements in place with two brokers, each
of which allows you access to their API to obtain prices:

- Whizzy FX, who return prices quickly but don't always offer the best rates.
- Steady FX, who don't return prices so quickly but tend to offer better rates.

The application your company uses to get prices and make choices isn't very good. It queries each broker sequentially,
and picks the best price. However, this sometimes takes so long that by the time you want to make a transaction, the
quoted price is no longer available.

## Exercise

Write two new BrokerSelectors which will perform concurrent queries:

- `SelectBest`, which waits until all prices are in and picks the best one.
- `SelectFastest`, which accepts the first price it receives.

Skeletons have been provided in the `exhcange/best` and `exchange/fastest` packages. The current selector implementation
is in `exchange/simple`.

## Hints

You may find the following things useful:

- Goroutines
- Channels
- The `sync.WaitGroup` type
- The `sync.Mutex` type

You should write unit tests to check the correctness of your selectors. If you want to run against test broker services,
run `server/run.sh`, which will start up two HTTP services for Whizzy FX and Speedy FX on localhost ports 8080 and 9090.
The `main.go` program is a test harness which you can use to create a selector and see its result.

## Reference examples

Reference implementations of the new selectors can be found on the `reference` branch of this repository. They represent
just two of the ways the problem could be solved and shouldn't be seen as the absolute correct answer.

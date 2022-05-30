package exchange

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
    "strings"
)

type Broker struct {
    Name string
    URL  string
}

func GetPrice(url, sell, buy string) (float64, error) {
    response, err := http.Get(url + fmt.Sprintf("/rate?sell=%s&buy=%s", sell, buy))
    if err != nil {
        return 0, err
    }

    defer response.Body.Close()
    body, err := io.ReadAll(response.Body)
    if err != nil {
        return 0, err
    }
    return strconv.ParseFloat(strings.TrimSpace(string(body)), 64)
}

type BrokerAPI func(url, sell, buy string) (float64, error)

type BrokerSelector interface {
    Select(brokers []Broker, sell, buy string) (float64, Broker)
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var baseDelay, delayRange int

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var err error
	baseDelay, err = strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	delayRange, err = strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/rate", rate)
	http.ListenAndServe(":" + os.Args[3], nil)
}

func rate(writer http.ResponseWriter, request *http.Request) {
	sell := request.URL.Query().Get("sell")
	buy := request.URL.Query().Get("buy")
	if len(sell) != 3 || len(buy) != 3 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	delayMillis := time.Duration(baseDelay + rand.Intn(delayRange))
	time.Sleep(time.Millisecond * delayMillis)

	rateStr := fmt.Sprintf("%f\n", rand.Float64())
	_, _ = writer.Write([]byte(rateStr))
}

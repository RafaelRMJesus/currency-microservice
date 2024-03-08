package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/RafaelRMJesus/currency-microservice/client"
)

func main() {
	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "BTC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", price)
	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	return
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr, svc)
	server.Run()
}

package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONApiServer(*listenAddr, svc)
	server.Run()
}

package elering

import (
	"fmt"
	"log"
	"testing"
)

func TestCall(t *testing.T) {

	fetcher := NewEleringFetcher("https://dashboard.elering.ee")
	resp, err := fetcher.NextDay()

	if err != nil {
		log.Panic(err.Error())
	}

	for _, price := range resp.Prices {
		fmt.Printf("%s - %s - %.2f \n", price.TimeUTC, price.TimeEestiAeg, price.Price)
	}
}

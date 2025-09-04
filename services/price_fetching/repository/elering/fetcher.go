package elering

import (
	"fmt"
	"log"
	"time"

	"resty.dev/v3"
)

type EleringFetcher struct {
	dns string
}

const (
	CSVResource = "api/nps/price/csv"
)

func NewEleringFetcher(dns string) *EleringFetcher {
	return &EleringFetcher{
		dns: dns,
	}
}

func (fetcher *EleringFetcher) FetchNextDay() {
	client := resty.New()
	defer client.Close()

	now := time.Now()

	x := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	begining := x.AddDate(0, 0, 1).Add(-3 * time.Hour)
	end := x.AddDate(0, 0, 2).Add(-3 * time.Hour)

	log.Printf("Query Params: \n \t\tBegining:%s \n \t\tEnd:%s ", begining.Format(time.RFC3339), end.Format(time.RFC3339))

	res, err := client.R().
		EnableTrace().
		SetQueryParam("start", begining.Format(time.RFC3339)).
		SetQueryParam("end", end.Format(time.RFC3339)).
		SetQueryParam("fields", "ee").
		Get(fmt.Sprintf("%s/%s", fetcher.dns, CSVResource))

	log.Println(err, res)
	log.Println(res.Request.TraceInfo())
}

package elering

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/core/domain"
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

func (fetcher *EleringFetcher) NextDay() (*domain.EnergyPrices, error) {
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

	prices, err := fetcher.readCSV(res)

	if err != nil {
		log.Println("err while saving the csv into struct, err: ", err.Error())
		return nil, err
	}

	return &domain.EnergyPrices{
		From:    begining,
		To:      end,
		TakenOn: now,
		Prices:  prices,
	}, nil
}

func (fetcher *EleringFetcher) readCSV(resp *resty.Response) ([]*domain.EnergyPrice, error) {
	csvData := resp.String()
	reader := csv.NewReader(strings.NewReader(csvData))
	reader.Comma = ';'

	var prices []*domain.EnergyPrice
	_, err := reader.Read()
	if err != nil {
		log.Println("err while skiping the header", err.Error())
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("err while reading", err.Error())
			return nil, err
		}

		price, err := strconv.ParseFloat(strings.ReplaceAll(record[2], ",", "."), 32)

		if err != nil {
			log.Println("err while transforming to float", err.Error())
			return nil, err
		}

		prices = append(prices, &domain.EnergyPrice{
			TimeUTC:      record[0],
			TimeEestiAeg: record[1],
			Price:        float32(price),
		})
	}

	return prices, nil
}

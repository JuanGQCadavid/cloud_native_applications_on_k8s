package main

import (
	"log"

	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/core"
	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/repository/elering"
	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/repository/local"
)

var service *core.Service

func init() {
	service = core.NewService(
		elering.NewEleringFetcher("https://dashboard.elering.ee"),
		local.NewCSVSaver("/tmp/price-data", "prices.csv"),
	)
}

func main() {
	if err := service.Run(); err != nil {
		log.Println("Cataplun! ", err.Error())
		return
	}

	log.Println("Done!")
}

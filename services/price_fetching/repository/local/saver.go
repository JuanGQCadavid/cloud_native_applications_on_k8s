package local

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/core/domain"
)

type CSVSaver struct {
	path     string
	fileName string
}

func NewCSVSaver(path, fileName string) *CSVSaver {
	return &CSVSaver{
		path:     path,
		fileName: fileName,
	}
}

func (saver *CSVSaver) Save(prices *domain.EnergyPrices) error {
	filePath := fmt.Sprintf("%s/%s", saver.path, saver.fileName)
	log.Println("Saving file on: ", filePath)
	file, err := os.Create(filePath)

	if err != nil {
		log.Println("Error while creating the file: ", err.Error())
		return nil
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Ajatempel (UTC)", "Eesti aeg", "NPS Eesti"})

	for _, r := range prices.Prices {
		row := []string{
			r.TimeUTC,
			r.TimeEestiAeg,
			fmt.Sprintf("%.2f", r.Price),
		}
		if err := writer.Write(row); err != nil {
			log.Println("Error writing record:", err)
		}
	}

	log.Println("Saved!")

	return nil
}

package core

import (
	"log"

	"github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/core/ports"
)

type Service struct {
	fetcher ports.Fetcher
	saver   ports.Saver
}

func NewService(fetcher ports.Fetcher, saver ports.Saver) *Service {
	return &Service{
		fetcher: fetcher,
		saver:   saver,
	}
}

func (srv *Service) Run() error {
	data, err := srv.fetcher.NextDay()

	if err != nil {
		log.Println("Fetcher error")
		return err
	}

	if err := srv.saver.Save(data); err != nil {
		log.Println("Saver error")
		return err
	}

	return nil
}

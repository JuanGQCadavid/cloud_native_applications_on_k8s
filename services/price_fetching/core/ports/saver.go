package ports

import "github.com/JuanGQCadavid/cloud_native_applications_on_k8s/price_fetching/core/domain"

type Saver interface {
	Save(*domain.EnergyPrices) error
}

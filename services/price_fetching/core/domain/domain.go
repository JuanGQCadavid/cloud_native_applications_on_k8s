package domain

import "time"

type EnergyPrice struct {
	TimeUTC      string
	TimeEestiAeg string
	Price        float32
}

type EnergyPrices struct {
	From    time.Time
	To      time.Time
	TakenOn time.Time
	Prices  []*EnergyPrice
}

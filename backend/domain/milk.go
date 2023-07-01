package domain

import "time"

type Milk struct {
	Date   time.Time
	Liters float64
	Price  float64
}

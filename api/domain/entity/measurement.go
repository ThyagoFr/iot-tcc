package entity

import (
	"time"
)

type Measurement struct {
	Entity
	DeviceID    int
	Temperature float64
	Turbidity   float64
	PH          float64
	Timestamp   time.Time
}

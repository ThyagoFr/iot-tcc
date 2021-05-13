package measurement

import (
	"time"
)

type Measurement struct {
	DeviceID    int                `bson:"deviceID" json:"deviceID"`
	Temperature float64            `bson:"temperature" json:"temperature"`
	Turbidity   float64            `bson:"turbidity" json:"turbidity"`
	PH          float64            `bson:"ph" json:"ph"`
	Timestamp   time.Time          `bson:"timestamp" json:"timestamp"`
}

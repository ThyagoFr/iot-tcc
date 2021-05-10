package device

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceData struct {
	ID          primitive.ObjectID `bson:"_id"`
	DeviceID    int                `bson:"device_id"`
	Temperature float64            `bson:"temperature"`
	Turbidity   float64            `bson:"turbidity"`
	PH          float64            `bson:"ph"`
	Timestamp   time.Time          `bson:"timestamp"`
}

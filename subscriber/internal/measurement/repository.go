package measurement

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, data *Measurement) error
	InsertBatch(ctx context.Context, batch []*Measurement) error
	LastMeasurementsForDevice(ctx context.Context, deviceID int) (*Measurement,error)
	Historical(ctx context.Context, sensorID int, to time.Time, from time.Time) ([]*Measurement, error)
}

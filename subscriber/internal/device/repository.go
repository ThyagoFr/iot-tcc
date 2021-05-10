package device

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context,data *DeviceData) error
	InsertBatch(ctx context.Context,batch []*DeviceData) error
	LastInfo(ctx context.Context,sensorID int) (*DeviceData, error)
	Historical(ctx context.Context,sensorID int, to time.Time, from time.Time) ([]*DeviceData, error)
}
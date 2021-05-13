package device

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, data *Device) error
	InsertBatch(ctx context.Context, batch []*Device) error
	GetAll(ctx context.Context) ([]*Device, error)
}

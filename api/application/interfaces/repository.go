package interfaces

import (
	"context"
	"time"

	"github.com/thyago/tcc/api-service/domain/entity"
)

type DeviceRepository interface {
	Insert(ctx context.Context, device *entity.Device) error
	InsertBatch(ctx context.Context, batch []*entity.Device) error
	FindAll(ctx context.Context) ([]*entity.Device, error)
}

type MeasurementRepository interface {
	Insert(ctx context.Context, measurement *entity.Measurement) error
	InsertBatch(ctx context.Context, batch []*entity.Measurement) error
	FindLastByDeviceID(ctx context.Context, deviceID int) (*entity.Measurement, error)
	FindByDeviceIDRangeDate(ctx context.Context, deviceID string, to time.Time, from time.Time) ([]*entity.Measurement, error)
}

type NotificationRepository interface {
	Insert(ctx context.Context, notification *entity.Notification) error
	Update(ctx context.Context, notification *entity.Notification) error
	FindAll() ([]*entity.Notification, error)
	FindByDeviceID(ctx context.Context, deviceID string) ([]*entity.Notification, error)
	Remove(ctx context.Context, notificationID string) error
}

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, userID string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Remove(ctx context.Context, userID string) error
	Update(ctx context.Context, user *entity.User) error
}

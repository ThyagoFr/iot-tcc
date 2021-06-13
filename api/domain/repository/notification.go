package repository

import (
	"context"
	"github.com/thyagofr/tcc/api/domain/entity"
)

type NotificationRepository interface {
	Insert(ctx context.Context, notification *entity.Notification) error
	Update(ctx context.Context, notification *entity.Notification) error
	FindAll() ([]*entity.Notification, error)
	FindByDeviceID(ctx context.Context, deviceID string) ([]*entity.Notification, error)
	Remove(ctx context.Context, notificationID string) error
}
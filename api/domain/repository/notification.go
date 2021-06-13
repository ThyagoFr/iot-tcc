package repository

import (
	"github.com/thyagofr/tcc/api/domain/entity"
)

type NotificationRepository interface {
	Insert(notification *entity.Notification) error
	Update(notification *entity.Notification) error
	FindAll() ([]*entity.Notification, error)
	FindByDeviceID(deviceID string) ([]*entity.Notification, error)
	Remove(notificationID string) error
}

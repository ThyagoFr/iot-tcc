package repositories

import (
	"github.com/thyagofr/tcc/api/domain/entity"
	"github.com/thyagofr/tcc/api/domain/repository"
	"gorm.io/gorm"
)

type NotificationPostgresql struct {
	db *gorm.DB
}

func NewNotificationPostgresql(db *gorm.DB) repository.NotificationRepository {
	return &NotificationPostgresql{
		db: db,
	}
}

func (pg *NotificationPostgresql) Insert(notification *entity.Notification) error {
	return nil
}

func (pg *NotificationPostgresql) Update(notification *entity.Notification) error {
	return nil
}

func (pg *NotificationPostgresql) FindAll() ([]*entity.Notification, error) {
	return nil, nil
}

func (pg *NotificationPostgresql) FindByDeviceID(deviceID string) ([]*entity.Notification, error) {
	return nil, nil
}

func (pg *NotificationPostgresql) Remove(notificationID string) error {
	return nil
}

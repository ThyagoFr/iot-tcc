package repositories

import (
	"context"

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

func (pg *NotificationPostgresql) Insert(ctx context.Context, notification *entity.Notification) error {
	return nil
}

func (pg *NotificationPostgresql) Update(ctx context.Context, notification *entity.Notification) error {
	return nil
}

func (pg *NotificationPostgresql) FindAll() ([]*entity.Notification, error) {
	return nil, nil
}

func (pg *NotificationPostgresql) FindByDeviceID(ctx context.Context, deviceID string) ([]*entity.Notification, error) {
	return nil, nil
}

func (pg *NotificationPostgresql) Remove(ctx context.Context, notificationID string) error {
	return nil
}

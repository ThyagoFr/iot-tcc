package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/thyago/tcc/api-service/application/interfaces"
	"github.com/thyago/tcc/api-service/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationMongoDB struct {
	collection *mongo.Collection
}

func NewNotificationMongoDB(db *mongo.Database) interfaces.NotificationRepository {
	return &NotificationMongoDB{
		collection: db.Collection("notifications"),
	}
}

func (n *NotificationMongoDB) Insert(ctx context.Context, notification *entity.Notification) error {
	notification.BeforeSave()
	_, err := n.collection.InsertOne(ctx, notification)
	return err
}

func (n *NotificationMongoDB) Update(ctx context.Context, notification *entity.Notification) error {

}

func (n *NotificationMongoDB) FindAll() ([]*entity.Notification, error) {

}

func (n *NotificationMongoDB) FindByDeviceID(ctx context.Context, deviceID string) ([]*entity.Notification, error) {

}

func (n *NotificationMongoDB) Remove(ctx context.Context, notificationID string) error {
	update := bson.D{{"$set", bson.D{{"deleted_at", time.Now()}}}}
	res, err := n.collection.UpdateOne(ctx, bson.D{{"ID", notificationID}}, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("notification not found")
	}
	return nil
}

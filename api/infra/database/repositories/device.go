package repositories

import (
	"context"

	"github.com/thyago/tcc/api-service/application/interfaces"
	"github.com/thyago/tcc/api-service/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceMongoDB struct {
	collection *mongo.Collection
}

func NewDeviceMongoDB(db *mongo.Database) interfaces.DeviceRepository {
	return &DeviceMongoDB{
		collection: db.Collection("devices"),
	}
}

func (mdb *DeviceMongoDB) Insert(ctx context.Context, device *entity.Device) error {
	device.BeforeSave()
	_, err := mdb.collection.InsertOne(ctx, device)
	return err
}

func (mdb *DeviceMongoDB) InsertBatch(ctx context.Context, batch []*entity.Device) error {
	dataToInsert := make([]interface{}, len(batch))
	for index, data := range batch {
		data.BeforeSave()
		dataToInsert[index] = data
	}
	_, err := mdb.collection.InsertMany(ctx, dataToInsert)
	return err
}

func (mdb *DeviceMongoDB) FindAll(ctx context.Context) ([]*entity.Device, error) {
	cursor, err := mdb.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	devices := make([]*entity.Device, 0)
	err = cursor.All(ctx, &devices)
	return devices, err
}

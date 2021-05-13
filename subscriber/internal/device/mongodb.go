package device

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongoDB(db *mongo.Database) Repository {
	return &MongoDB{
		collection: db.Collection("devices"),
	}
}

func (mdb *MongoDB) Insert(ctx context.Context, device *Device) error {
	_, err := mdb.collection.InsertOne(ctx, device)
	return err
}

func (mdb *MongoDB) InsertBatch(ctx context.Context, batch []*Device) error {
	dataToInsert := make([]interface{}, len(batch))
	for index, data := range batch {
		dataToInsert[index] = data
	}
	_, err := mdb.collection.InsertMany(ctx, dataToInsert)
	return err
}

func (mdb *MongoDB) GetAll(ctx context.Context) ([]*Device, error) {
	cursor, err := mdb.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil,err
	}
	defer cursor.Close(ctx)
	devices := make([]*Device, 0)
	err = cursor.All(ctx, &devices)
	return devices, err
}

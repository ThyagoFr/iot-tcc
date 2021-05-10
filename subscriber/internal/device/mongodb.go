package device

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongoDB(db *mongo.Database) Repository {
	return &MongoDB{
		collection: db.Collection("device_data"),
	}
}

func (mdb *MongoDB) Insert(ctx context.Context,data *DeviceData) error {
	_, err := mdb.collection.InsertOne(ctx, data)
	return err
}

func (mdb *MongoDB) InsertBatch(ctx context.Context, batch []*DeviceData) error {
	dataToInsert := make([]interface{}, len(batch))
	for index, data := range batch {
		dataToInsert[index] = data
	}
	_, err := mdb.collection.InsertMany(ctx, dataToInsert)
	return err
}

func (mdb *MongoDB) LastInfo(ctx context.Context,sensorID int) (*DeviceData, error) {
	data := &DeviceData{}
	err := mdb.collection.FindOne(
		ctx,
		bson.M{
		"sensor_id" : sensorID,
		},
		options.FindOne().SetSort(bson.D{{"_id", -1}}),
	).Decode(data)
	return data, err
}

func (mdb *MongoDB) Historical(ctx context.Context,sensorID int, to time.Time, from time.Time) ([]*DeviceData, error) {
	data := make([]*DeviceData,0)
	cursor, err := mdb.collection.Find(
		ctx,
		bson.M{
			"sensor_id" : sensorID,
			"timestamp" : bson.M{
				"$gte" : from,
				"$lt" : to,
			},
		},
	)
	if err != nil {
		return nil,err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx,&data)
	if err != nil {
		return nil,err
	}
	return data,nil
}




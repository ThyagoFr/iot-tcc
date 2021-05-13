package measurement

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
		collection: db.Collection("measurements"),
	}
}

func (mdb *MongoDB) Insert(ctx context.Context, data *Measurement) error {
	_, err := mdb.collection.InsertOne(ctx, data)
	return err
}

func (mdb *MongoDB) InsertBatch(ctx context.Context, batch []*Measurement) error {
	dataToInsert := make([]interface{}, len(batch))
	for index, data := range batch {
		dataToInsert[index] = data
	}
	_, err := mdb.collection.InsertMany(ctx, dataToInsert)
	return err
}

func (mdb *MongoDB) LastMeasurementsForDevice(ctx context.Context, sensorID int) (*Measurement, error) {
	data := &Measurement{}
	err := mdb.collection.FindOne(
		ctx,
		bson.M{
			"deviceID": sensorID,
		},
		options.FindOne().SetSort(bson.D{{"_id", -1}}),
	).Decode(data)
	return data, err
}

func (mdb *MongoDB) Historical(ctx context.Context, deviceID int, to time.Time, from time.Time) ([]*Measurement, error) {
	data := make([]*Measurement, 0)
	cursor, err := mdb.collection.Find(
		ctx,
		bson.M{
			"deviceID": deviceID,
			"timestamp": bson.M{
				"$gte": from,
				"$lt":  to,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

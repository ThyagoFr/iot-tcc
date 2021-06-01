package repositories

import (
	"context"
	"time"

	"github.com/thyago/tcc/api-service/application/interfaces"
	"github.com/thyago/tcc/api-service/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MeasurementMongoDB struct {
	collection *mongo.Collection
}

func NewMeasurementMongoDB(db *mongo.Database) interfaces.MeasurementRepository {
	return &MeasurementMongoDB{
		collection: db.Collection("measurements"),
	}
}

func (mdb *MeasurementMongoDB) Insert(ctx context.Context, data *entity.Measurement) error {
	data.BeforeSave()
	_, err := mdb.collection.InsertOne(ctx, data)
	return err
}

func (mdb *MeasurementMongoDB) InsertBatch(ctx context.Context, batch []*entity.Measurement) error {
	dataToInsert := make([]interface{}, len(batch))
	for index, data := range batch {
		data.BeforeSave()
		dataToInsert[index] = data
	}
	_, err := mdb.collection.InsertMany(ctx, dataToInsert)
	return err
}

func (mdb *MeasurementMongoDB) FindLastByDeviceID(ctx context.Context, sensorID int) (*entity.Measurement, error) {
	data := &entity.Measurement{}
	err := mdb.collection.FindOne(
		ctx,
		bson.M{
			"deviceID": sensorID,
		},
		options.FindOne().SetSort(bson.D{{"_id", -1}}),
	).Decode(data)
	return data, err
}

func (mdb *MeasurementMongoDB) FindByDeviceIDRangeDate(ctx context.Context, deviceID string, to time.Time, from time.Time) ([]*entity.Measurement, error) {
	data := make([]*entity.Measurement, 0)
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

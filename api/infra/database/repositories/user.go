package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/thyago/tcc/api-service/application/interfaces"
	"github.com/thyago/tcc/api-service/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoDB struct {
	collection *mongo.Collection
}

func NewUserMongoDB(db *mongo.Database) interfaces.UserRepository {
	return &UserMongoDB{
		collection: db.Collection("users"),
	}
}

func (u *UserMongoDB) Insert(ctx context.Context, user *entity.User) error {
	user.BeforeSave()
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u *UserMongoDB) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	user := &entity.User{}
	err := u.collection.FindOne(
		ctx,
		bson.M{
			"ID": userID,
		},
	).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (u *UserMongoDB) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.collection.FindOne(
		ctx,
		bson.M{
			"email": email,
		},
	).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (u *UserMongoDB) FindAll(ctx context.Context) ([]*entity.User, error) {
	cursor, err := u.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	users := make([]*entity.User, 0)
	err = cursor.All(ctx, &users)
	return users, err
}

func (u *UserMongoDB) Remove(ctx context.Context, userID string) error {
	update := bson.D{{"$set", bson.D{{"deleted_at", time.Now()}}}}
	res, err := u.collection.UpdateOne(ctx, bson.D{{"ID", userID}}, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserMongoDB) Update(ctx context.Context, user *entity.User) error {
	update := bson.D{{"$set",
		bson.D{
			{"updated_at", time.Now()},
			{"name", user.Name},
			{"email", user.Email},
		},
	},
	}
	res, err := u.collection.UpdateOne(ctx, bson.D{{"ID", user.ID}}, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("user not found")
	}
	return nil

}

package mongo

import (
	"context"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificationMongo struct {
	db *mongo.Client
}

func NewVerificationMongo(db *mongo.Client) *VerificationMongo {
	return &VerificationMongo{db: db}
}

func (r *VerificationMongo) SetVerificationCode(id string, code int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "lastVerificationCode", Value: code},
		}},
	}

	_, err := r.db.Database("db").Collection("users").UpdateByID(ctx, id, update)

	return err
}

func (r *VerificationMongo) GetVerificationCode(id string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User

	err := r.db.Database("db").Collection("users").FindOne(ctx, bson.D{
		{Key: "_id", Value: id},
	}).Decode(&user)

	return user.LastVerificationCode, err
}

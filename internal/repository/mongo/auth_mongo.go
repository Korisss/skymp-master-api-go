package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if r.db.Database("db").Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Err() == nil {
		return "", errors.New("user already registered")
	}

	insertResult, err := r.db.Database("db").Collection("users").InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AuthMongo) GetUser(email, password string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User

	err := r.db.Database("db").Collection("users").FindOne(ctx, bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}).Decode(&user)

	return user, err
}

func (r *AuthMongo) GetUserName(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorln("Invalid id")
	}

	var user domain.User

	err = r.db.Database("db").Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)

	return user.Name, err
}

package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user domain.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if r.db.Database("db").Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Err() == nil {
		return -1, errors.New("user already registered")
	}

	id, err := r.db.Database("db").Collection("users").CountDocuments(ctx, bson.M{})
	if err != nil {
		return -1, err
	}

	user.Id = id + 1
	_, err = r.db.Database("db").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return -1, err
	}

	return user.Id, nil
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

func (r *AuthMongo) GetUserName(id int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User

	err := r.db.Database("db").Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(&user)

	return user.Name, err
}

package mongo

// import (
// 	"context"
// 	"time"

// 	"github.com/Korisss/skymp-master-api-go/internal/domain"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type VerificationMongo struct {
// 	db *mongo.Client
// }

// func NewVerificationMongo(db *mongo.Client) *VerificationMongo {
// 	return &VerificationMongo{db: db}
// }

// func (r *VerificationMongo) SetDiscord(id int64, discord string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	update := bson.D{
// 		{Key: "$set", Value: bson.D{
// 			{Key: "discord", Value: discord},
// 		}},
// 	}

// 	_, err := r.db.Database("db").Collection("users").UpdateByID(ctx, id, update)

// 	return err
// }

// func (r *VerificationMongo) SetVerificationCode(id int64, code int) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	update := bson.D{
// 		{Key: "$set", Value: bson.D{
// 			{Key: "lastVerificationCode", Value: code},
// 		}},
// 	}

// 	_, err := r.db.Database("db").Collection("users").UpdateByID(ctx, id, update)

// 	return err
// }

// func (r *VerificationMongo) GetVerificationCode(id int64) (int, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var user domain.User

// 	err := r.db.Database("db").Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(&user)

// 	return user.LastVerificationCode, err
// }

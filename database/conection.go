package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() *mongo.Client {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test_user:testpassword123@cluster0.577kp.mongodb.net/covid?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, _ := GetDbContext()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client
}

func DisconnectDb(client *mongo.Client) {
	ctx, _ := GetDbContext()
	defer client.Disconnect(ctx)
}

func InsertToDb(client *mongo.Client, docs []interface{}) error {
	ctx, _ := GetDbContext()
	collection := client.Database("covid").Collection("col1")
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}

	_, err = collection.InsertMany(ctx, docs)
	return err
}

func FindDocumentFromState(client *mongo.Client, state string) *mongo.Cursor {
	ctx, _ := GetDbContext()
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"state", state}},
			bson.D{{"state", "India"}},
		}},
	}
	option := bson.D{{"_id", 0}}
	collection := client.Database("covid").Collection("col1")
	result, err := collection.Find(ctx, filter, options.Find().SetProjection(option))
	if err != nil {
		panic(err)
	}
	return result
}

func GetDbContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://tanmoysantra67:<password>@cluster0.ozd0zhu.mongodb.net/?retryWrites=true&w=majority"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mogodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

}

// MONGODB helpers - file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ",inserted.InsertedID)

}

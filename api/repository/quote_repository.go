package quoteRepository

import (
	"context"
	"fmt"
	"log"

	"github.com/cbellee/shutter-quote-app/api"
	"github.com/cbellee/shutter-quote-app/api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	conf, err    = config.LoadConfig()
	dbName       = "quoteAppDB"
	dbCollection = "quotes"
)

// QuoteRepository used to get quote data from mongodDB
type QuoteRepository interface {
	Get(id int64) (*api.Quote, error)
	List() ([]*api.Quote, error)
	Insert(quote api.Quote) (lastInsertID interface{}, err error)
	Delete(id int64) error
	Update(quote api.Quote, id int64) (upsertedCount int64, err error)
}

type quoteRepository struct {
	client *mongo.Client
}

// NewQuoteRepository returns a new instance of QuoteRepository
func NewQuoteRepository(client *mongo.Client) QuoteRepository {
	return &quoteRepository{
		client: client,
	}
}

// Get
func (r *quoteRepository) Get(id int64) (quote *api.Quote, err error) {
	var result *api.Quote
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// List
func (r *quoteRepository) List() (quotes []*api.Quote, err error) {
	var results []*api.Quote
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	//defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var element api.Quote
		err := cur.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &element)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	fmt.Printf("Found %d documents\n", len(results))
	return results, nil
}

// Delete
func (r *quoteRepository) Delete(id int64) (err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v document in '%s' collection\n", result.DeletedCount, conf.DbCollection)
	return err
}

// Insert
func (r *quoteRepository) Insert(quote api.Quote) (lastInsertID interface{}, err error) {
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	insertResult, err := collection.InsertOne(context.Background(), quote)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID, err
}

// Update
func (r *quoteRepository) Update(quote api.Quote, id int64) (upsertedCount int64, err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	updateResult, err := collection.UpdateOne(context.Background(), filter, quote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents in '%s' collection\n", updateResult.UpsertedCount, conf.DbCollection)
	return updateResult.UpsertedCount, err
}

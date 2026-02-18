package caching

import (
	"arbolito/pkg/model"
	"arbolito/pkg/repository"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "rates"
)

type mongoCachingRepository struct {
	db *mongo.Database
}

func NewMongoCachingRepository(db *mongo.Database) (repository.CachingRepository, error) {
	repo := &mongoCachingRepository{db: db}
	err := repo.createTTLIndex()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *mongoCachingRepository) GetRate() (*model.CachedRate, error) {
	var cachedRate model.CachedRate
	collection := r.db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find the most recent entry
	opts := options.FindOne().SetSort(bson.D{{"created_at", -1}})
	err := collection.FindOne(ctx, bson.D{}, opts).Decode(&cachedRate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Cache miss: no rate found in cache")
			return nil, nil // No documents found is not an error here
		}
		log.Printf("Error retrieving rate from cache: %v", err)
		return nil, err
	}
	log.Printf("Cache hit: retrieved rate from cache (created at %v)", cachedRate.CreatedAt)
	return &cachedRate, nil
}

func (r *mongoCachingRepository) SetRate(rate *model.Rate) error {
	log.Printf("Setting rate in cache")
	collection := r.db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cachedRate := model.CachedRate{
		Rate:      *rate,
		CreatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, cachedRate)
	if err != nil {
		log.Printf("Error setting rate in cache: %v", err)
		return err
	}
	log.Printf("Successfully set rate in cache")
	return nil
}

func (r *mongoCachingRepository) createTTLIndex() error {
	collection := r.db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"created_at", 1}},
		Options: options.Index().SetExpireAfterSeconds(15 * 60), // 15 minutes
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

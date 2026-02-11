package mongodb

import (
	"context"
	"go-bookstore-mux-mongodb/internal/category/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BookstoreRepository struct {
	repo *mongo.Collection
}

func TitleUniqueIndex(coll *mongo.Collection) error {
	model := mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(context.Background(), model)
	return err
}

func NewBookstoreRepository(col *mongo.Collection) *BookstoreRepository {
	return &BookstoreRepository{
		repo: col,
	}
}

func (r *BookstoreRepository) Add(b domain.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"book": bson.M{
			"title":       b.Title,
			"description": b.Description,
			"author":      b.Author,
			"value":       b.Value,
			"release":     b.ReleaseDate,
		},
	}

	if _, err := r.repo.InsertOne(ctx, doc); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return err
		}
		return err
	}

	return nil
}

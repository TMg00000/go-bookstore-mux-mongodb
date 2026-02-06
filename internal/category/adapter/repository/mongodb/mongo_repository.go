package mongodb

import (
	"context"
	"go-bookstore-mux-mongodb/internal/category/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BookstoreRepository struct {
	repo *mongo.Collection
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
		return err
	}

	return nil
}

func (r *BookstoreRepository) Search(title string) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book domain.Book

	filter := bson.M{"title": title}
	err := r.repo.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil

}

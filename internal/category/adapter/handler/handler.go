package handler

import (
	"encoding/json"
	"go-bookstore-mux-mongodb/internal/category/domain"
	"go-bookstore-mux-mongodb/internal/validation"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServicesRepository struct {
	Services domain.ServicesMongoDBRepository
}

type requestAddBooks struct {
	Title       string  `json:"title" bson:"title" validate:"required,min=2,max=30"`
	Description string  `json:"description" bson:"description" validate:"required,min=30,max=1500"`
	ReleaseDate string  `json:"release" bson:"release" validate:"required,datetime=2006-01-02"`
	Value       float32 `json:"value" bson:"value" validate:"required,min=1"`
	Author      string  `json:"author" bson:"author" validate:"required,min=3,max=50"`
}
type responseAddBooks struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	ReleaseDate time.Time          `json:"release" bson:"release"`
	Value       float32            `json:"value" bson:"value"`
	Author      string             `json:"author" bson:"author"`
}

var reqAddBook requestAddBooks
var respAddBook responseAddBooks

func (s *ServicesRepository) AddNewBook(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&reqAddBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validation.ValidateRequestStruct(reqAddBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook := domain.BookTypePost{
		Title:       reqAddBook.Title,
		Description: reqAddBook.Description,
		Author:      reqAddBook.Author,
		Value:       reqAddBook.Value,
		ReleaseDate: reqAddBook.ReleaseDate,
	}

	if err := s.Services.Add(newBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respAddBook)

}

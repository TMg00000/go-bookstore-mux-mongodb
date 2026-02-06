package main

import (
	"context"
	"go-bookstore-mux-mongodb/internal/category/adapter/handler"
	"go-bookstore-mux-mongodb/internal/category/adapter/repository/mongodb"
	"go-bookstore-mux-mongodb/internal/configs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	errConfig := configs.StartConfig()
	returnFatalError(errConfig)

	db := initMongoDB()
	defer func() {
		err := db.Database().Client().Disconnect(context.Background())
		returnFatalError(err)
	}()

	bookstoreRepo := mongodb.NewBookstoreRepository(db)
	router := handler.Repository{
		Repo: bookstoreRepo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/bookstore", router.AddNewBook).Methods("POST")

	log.Printf("Rodando na porta :8080")
	log.Printf("Rodando no endereço http://localhost:8080")

	returnFatalError(http.ListenAndServe(":8080", r))
}

func initMongoDB() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongodb.NewMongoConnection(ctx)
	returnFatalError(err)
	configs.LoadEnv(client)

	return client.Database(configs.Env.DbName).Collection(configs.Env.ColName)
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

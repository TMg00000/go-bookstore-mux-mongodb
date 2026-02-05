package configs

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type config struct {
	MongoURI string `envconfig:"MONGO_URI"`
	DbName   string `envconfig:"DB_NAME"`
	ColName  string `envconfig:"COLLECTION_NAME"`
	Port     int32  `envconfig:"PORT"`
}

var Env config

func StartConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	if Env.MongoURI == "" {
		log.Fatal("MONGO_URI obrigatório")
	}

	return nil
}

func LoadEnv(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := client.Database("bookstore").Collection("configs")

	var result bson.M
	err := col.FindOne(ctx, bson.M{"_id": "app_config"}).Decode(&result)
	if err != nil {
		panic(err)
	}

	Env.DbName = result["DB_NAME"].(string)
	Env.ColName = result["COLLECTION_NAME"].(string)
	Env.Port = result["PORT"].(int32)
}

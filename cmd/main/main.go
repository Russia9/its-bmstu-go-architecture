package main

import (
	"context"
	"github.com/Russia9/its-bmstu-go-architecture/internal/post/delivery/rest"
	"github.com/Russia9/its-bmstu-go-architecture/internal/post/repository/mongodb"
	"github.com/Russia9/its-bmstu-go-architecture/internal/post/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {
	router := gin.Default()
	group := router.Group("/post")

	// MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	db := client.Database("blog")

	repo := mongodb.NewPostRepository(db)
	uc := usecase.NewPostUsecase(repo)

	rest.NewPostDelivery(uc, group)

	router.Run()
}

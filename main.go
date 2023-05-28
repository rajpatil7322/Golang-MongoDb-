package main

import (
	"context"
	"fmt"
	"mongo_golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:name", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:name", uc.DeleteUser)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mongo.Client {
	client, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("Got client")
	return client

}

func connect(uri string) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client, err
}

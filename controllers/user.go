package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"mongo_golang/models"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	fmt.Println(reflect.TypeOf(name))
	var filter interface{}

	filter = bson.D{
		{Key: "name", Value: &name},
	}

	res, err := uc.session.Database("mongo-golang").Collection("users").Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	var results []bson.D

	if err := res.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// printing the result of query.
	fmt.Println("Query Result")
	for _, doc := range results {
		fmt.Println(doc)
	}

	// uj, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	json.NewDecoder(r.Body).Decode(&u)

	// u.ID = bson.NewObjectId()

	// result, err := uc.session.Database("mongo-golang").Collection("users").InsertOne(context.TODO(), bson.D{
	// 	{Key: "name", Value: u.Name},
	// 	{Key: "gender", Value: u.Gender},
	// 	{Key: "age", Value: 20},
	// })

	uc.session.Database("mongo-golang").Collection("users").InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	var filter interface{}

	filter = bson.D{
		{Key: "name", Value: &name},
	}

	res, err := uc.session.Database("mongo-golang").Collection("users").DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}

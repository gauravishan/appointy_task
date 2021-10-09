package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	id   primitive.ObjectID `json: "_id"`
	name string             `json: "name"`
	// FieldInt int    `json: "Field Int"`
	// FieldBool  bool   `json: "Field Bool"`
	email    string `json: "email"`
	password string `json: "password"`
}

type Post struct {
	id      primitive.ObjectID `json: "_id"`
	caption string             `json: "caption"`
	// FieldInt int    `json: "Field Int"`
	// FieldBool  bool   `json: "Field Bool"`
	url  string `json: "url"`
	time string `json: "time"`
}

var client *mongo.Client

// func getHash(pwd []byte) string {
// 	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return string(hash)
// }

func userSignup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	user.id = primitive.NewObjectID()
	// user.password = getHash([]byte(user.password))
	// collection := client.Database("GODB").Collection("user")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// result, _ := collection.InsertOne(ctx, user)
	// json.NewEncoder(response).Encode(result)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("First_Database").Collection("First COllection")
	// fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := User{
		id:   user.id,
		name: user.name,
		// FieldBool:  true,
		email:    user.email,
		password: user.password,
	}

	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		// fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		// fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		// fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID

		// fmt.Fprintf(response, "+%v", user)

		fmt.Println("InsertedOne(), newID", newID)
		// fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hi")
// }

func handlerequest() {
	http.HandleFunc("/user", userSignup)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	// fmt.Println("Mongo.connect() ERROR: ", err)
	// 	os.Exit(1)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// col := client.Database("First_Database").Collection("First COllection")
	// // fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	// oneDoc := MongoField{
	// 	Fieldid:  "gaurav_ishan",
	// 	FieldStr: "Gaurav Ishan",
	// 	// FieldInt: 9334863893,
	// 	// FieldBool:  true,
	// 	Fieldemail:    "ishagaurav1904@gmail.com",
	// 	FieldPassword: "gaurav",
	// }

	// // fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	// result, insertErr := col.InsertOne(ctx, oneDoc)
	// if insertErr != nil {
	// 	// fmt.Println("InsertONE Error:", insertErr)
	// 	os.Exit(1)
	// } else {
	// 	// fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
	// 	// fmt.Println("InsertOne() api result type: ", result)

	// 	newID := result.InsertedID
	// 	fmt.Println("InsertedOne(), newID", newID)
	// 	// fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	// }

	handlerequest()
}

package main

import (
	"context"
	"encoding/json"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

	"fmt"
	"log"
	"net/http"

	//"context"
	//"encoding/json"
	"ngoc/api/winningnumbers/helper"
	"ngoc/api/winningnumbers/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func fuckingSayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `GO GO GO +_-`)
}

func getWinningNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var numbers models.WinningNumbers

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&numbers)

	// connect db
	collection := helper.ConnectDB()

	filter := bson.D{{"author", bson.D{{"firstname", "ngoc"}, {"lastname", "nguyen"}}}}

	var result models.WinningNumbers
	// insert our book model.
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func createWinningNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var numbers models.WinningNumbers

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&numbers)

	// connect db
	collection := helper.ConnectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), numbers)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	//Init Router
	r := mux.NewRouter()
	serverUp := getMessageServerUp()

	// arrange our route
	r.HandleFunc("/nan", getWinningNumbers).Methods("GET")
	r.HandleFunc("/nan", createWinningNumbers).Methods("POST")

	fmt.Println(serverUp)
	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMessageServerUp() string {
	return "Server up and running.."
}

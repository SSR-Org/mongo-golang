package main

import (
	"net/http" //creates a server with listen and serve function

	"github.com/julienschmidt/httprouter"
	"github.com/shravanth-drife/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()                             //new instance
	uc := controllers.NewUserController(getSession()) //creates new session
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r) //creates golang server at specified port
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107") //calls mongo port
	//error handling
	if err != nil {
		panic(err) //check statement for any error, and stops if there is any
	}
	return s
}

// func getSession() *mongo.Client {

// 	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// 	clientOptions := options.Client().
// 		ApplyURI("mongodb+srv://shravanth-drife:cEuh2ks5k34ACM8e@golang-api.jiu6c3y.mongodb.net/?retryWrites=true&w=majority").
// 		SetServerAPIOptions(serverAPIOptions)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	s, err := mongo.Connect(ctx, clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return s
// }

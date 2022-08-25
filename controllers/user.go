package controllers

import (
	"encoding/json" //golang doesnt understand json by default
	"fmt"
	"net/http" //creates a server with listen and serve function

	"github.com/julienschmidt/httprouter"
	"github.com/shravanth-drife/mongo-golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson" //Binary JSON library
)

type UserController struct {
	session *mgo.Session //struct with pointer as variable
}

func NewUserController(s *mgo.Session) *UserController { //returns pointer to user controller
	return &UserController{s} //address of User Controller
}

// w - Response back to Postman (200/404/..), r - Request from main.go, p - parameters
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //checks if id is hex or not, if not - 404
	}

	oid := bson.ObjectIdHex(id) //object id

	u := models.User{} //from models file, u stores the data from DB which we GET

	//creats DB if there is no existing DB, .C - collection of users
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // status 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{} //assigning the struct

	json.NewDecoder(r.Body).Decode(&u) //decode json values from postman, data caught in u

	u.Id = bson.NewObjectId() //creates new random object id

	uc.session.DB("mongo-golang").C("users").Insert(u) //insert new user

	uj, err := json.Marshal(u) //marshal back to json and send back to user
	if err != nil {
		fmt.Println(err)
	}

	//send response to front-end
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // status 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

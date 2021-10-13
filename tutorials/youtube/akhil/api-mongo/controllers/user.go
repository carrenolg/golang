package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/carrenolg/go-mongo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{
		session: s,
	}
}

// GetUser, CreateUser, DeleteUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// validation request
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	// user data
	objectId := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(objectId).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// encoding data
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", data)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}

	// decode request
	json.NewDecoder(r.Body).Decode(&user)

	// data
	user.Id = bson.NewObjectId()
	err := uc.session.DB("mongo-golang").C("users").Insert(user)
	if err != nil {
		log.Println(err)
	}

	// encoding data
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", data)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	// check
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	objectId := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(objectId); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", objectId, "\n")
}

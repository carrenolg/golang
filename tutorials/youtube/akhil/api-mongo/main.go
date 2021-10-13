package main

import (
	"net/http"

	"github.com/carrenolg/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// routes
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	// operations
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	strConn := "mongodb://root:admin@localhost:27017"
	session, err := mgo.Dial(strConn)
	if err != nil {
		panic(err)
	}
	return session
}

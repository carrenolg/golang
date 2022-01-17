package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Movie holds a movie data
type Movie struct {
	Name      string    `bson:"name"`
	Year      string    `bson:"year"`
	Directors []string  `bson:"directors"`
	Writers   []string  `bson:"writers"`
	BoxOffice BoxOffice `bson:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

const URL = "mongodb://root:admin@172.22.0.2:27017"

func main() {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Set up driver db
	c := session.DB("appdb").C("movies")
	// create a movie
	darkNight := &Movie{
		Name:      "The Dark Knight",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross:  533316061,
		},
	}
	// Insert into MongoDB
	err = c.Insert(darkNight)
	if err != nil {
		log.Fatal(err)
	}
	// Now query the movie back
	result := Movie{}
	// bson.M is used for nested fields
	err = c.Find(bson.M{
		"boxOffice.budget": bson.M{"$gt": 150000000},
	}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie:", result.Name)
}

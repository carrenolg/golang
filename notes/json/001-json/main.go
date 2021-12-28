package main

import (
	"encoding/json"
	"fmt"
	"log"

	"carrenolg.io/notes/json/models"
)

func main() {
	// 1. Ecoding
	m := models.Message{
		Name: "Alice",
		Body: "Hello",
		Time: 1294706395881547000,
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	// b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	fmt.Println(b)
	fmt.Println(string(b))

	//2. Decoding
	var v models.Message
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	//3. JSON vs. Go type
	dataJSON := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var dataGo models.Message
	// "Food" field doesn't match with the fields of Message type
	err = json.Unmarshal(dataJSON, &dataGo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataGo)
}

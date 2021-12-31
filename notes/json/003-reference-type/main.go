package main

import (
	"encoding/json"
	"fmt"
	"log"

	"carrenolg.io/notes/json/models"
)

type Map struct {
	Key string `json:"key"`
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var m models.FamilyMember
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	b = []byte(`{"Cmd":{"Type":1,"Text":"mkdir"}, "Msg": {}}`)
	var im models.IncomingMessage
	err = json.Unmarshal(b, &im)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*im.Cmd)
	fmt.Println(*im.Msg)

	input := []byte(`[{"key": "value"}]`)
	//var inventory []Size
	var inventory []interface{}
	if err := json.Unmarshal(input, &inventory); err != nil {
		log.Fatal(err)
	}
	fmt.Println(inventory[0])
	data := inventory[0].(map[string]interface{})
	fmt.Println(data["key"])

	// no interface
	var storage []Map
	if err := json.Unmarshal(input, &storage); err != nil {
		log.Fatal(err)
	}
	fmt.Println(storage)
}

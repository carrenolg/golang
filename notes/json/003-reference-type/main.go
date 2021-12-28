package main

import (
	"encoding/json"
	"fmt"
	"log"

	"carrenolg.io/notes/json/models"
)

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

}

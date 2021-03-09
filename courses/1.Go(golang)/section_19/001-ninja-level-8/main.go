package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	// Go to Json
	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Json data:", bs)

	fmt.Println("Json data:", string(bs))

	// exercise #2
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)

	var p []person

	bp := []byte(s)

	err = json.Unmarshal(bp, &p)

	//fmt.Println(p)

	for i, v := range p {
		fmt.Println(i, v)
	}

}

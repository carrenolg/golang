package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional requestOptions struct
	if err != nil {
		log.Fatal("Unable to make request:", err)
	}
	log.Println(resp.String())
	// Get data
	var returnData map[string]interface{}
	resp.JSON(&returnData)
	log.Println(returnData)
}

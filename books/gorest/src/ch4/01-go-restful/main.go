package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
)

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write to the response
	io.WriteString(resp, fmt.Sprintf("%s\n", time.Now()))
}

func main() {
	// Create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

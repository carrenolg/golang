package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/*GET takes a route and handler function*/
	r.GET("/pingTime", func(c *gin.Context) {
		// JSON serializer is avaliable on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000") // Listen an serve on 0.0.0.0:8000
}

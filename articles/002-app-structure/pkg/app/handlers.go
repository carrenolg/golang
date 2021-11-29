package app

import (
	"log"
	"net/http"

	"carrenolg.io/articles/app-structure/pkg/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

// add this below APIStatus method
func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newUser api.NewUserRequest

		err := c.ShouldBindJSON(&newUser)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.userService.New(newUser)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}

		c.JSON(http.StatusOK, response)
	}
}

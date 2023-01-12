package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Messsage string `json:"messsage"`
}

func newErrorResponse(c *gin.Context, code int, message string) {
	defer log.Fatalf("Error: %s", message)

	// stop event propagation (other handling) and send error message
	c.AbortWithStatusJSON(code, Error{message})
}

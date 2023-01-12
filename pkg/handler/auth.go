package handler

import (
	"github.com/gin-gonic/gin"
	todo_demo "github.com/semaffor/go-todo-app"
	"net/http"
)

// handler - is a function which should receive pointer on gin Context
func (h *Handler) logIn(c *gin.Context) {

}

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}

func (h *Handler) signUp(c *gin.Context) {
	var user todo_demo.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})
}

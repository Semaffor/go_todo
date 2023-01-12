package handler

import (
	"github.com/gin-gonic/gin"
	todo_demo "github.com/semaffor/go-todo-app"
	"net/http"
)

type AuthData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// handler - is a function which should receive pointer on gin Context
func (h *Handler) logIn(c *gin.Context) {
	var loginData AuthData

	if err := c.BindJSON(&loginData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(loginData.Username, loginData.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{"access-token": token})
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

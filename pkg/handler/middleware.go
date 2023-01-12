package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) identifyUser(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "auth header is empty")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "auth header is empty")
		return
	}

	//parse jwt
	userId, err := h.service.Authorization.ParseJwt(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	//Set id in context
	c.Set(userCtx, userId)
}

func getUserIdFromCtx(c *gin.Context) (int, error) {
	idStr, _ := c.Get(userCtx)
	id, ok := idStr.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid type: userId")
		return 0, errors.New("field userId hasn't found")
	}
	return id, nil
}

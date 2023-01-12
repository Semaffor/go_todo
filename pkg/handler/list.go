package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	todo_demo "github.com/semaffor/go-todo-app"
	"net/http"
	"strconv"
)

// @Summary Create list
// @Security JWT-token
// @Tags lists
// @Description create new empty list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Router /api/lists [post]

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserIdFromCtx(c)
	if err != nil {
		return
	}

	var list todo_demo.TodoList
	if err := c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listId, err := h.service.TodoList.Create(userId, list)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"listId": listId})
}

type AllUserLists struct {
	Data []todo_demo.TodoList
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserIdFromCtx(c)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &AllUserLists{Data: lists})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserIdFromCtx(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid id param")
		return
	}

	list, err := h.service.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserIdFromCtx(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo_demo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"result": "done"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserIdFromCtx(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid id param")
		return
	}

	code, err := h.service.TodoList.DeleteById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if code == 204 {
		answer := fmt.Sprintf("List with id=%d doesn't exist", listId)
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": answer})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"result": "done"})
	}
}

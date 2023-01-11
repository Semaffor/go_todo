package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/semaffor/go-todo-app/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// create route (similar in React)
	router := gin.New()

	// routes for auth
	auth := router.Group("/auth")
	{
		auth.POST("log-in", h.logIn)
		auth.POST("sign-up", h.signUp)
	}

	// routes (@RequestMapping("/api"))
	api := router.Group("/api")
	{
		// routes for todo list
		lists := api.Group("/lists")
		{
			// /api/lists/
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:listId", h.getListById)
			lists.DELETE("/:listId", h.updateList)
			lists.PUT("/:listId", h.deleteList)

			items := lists.Group(":listId/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:itemId", h.getItemById)
				items.DELETE("/:itemId", h.deleteItem)
				items.PUT("/:itemId", h.updateItem)
			}
		}
	}

	return router
}

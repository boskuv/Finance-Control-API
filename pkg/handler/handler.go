package handler

import (
	"github.com/gin-gonic/gin"
	//"github.com/zhashkevych/todo-app/pkg/service"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
//	services *service.Service
}

// func NewHandler(services *service.Service) *Handler {
// 	return &Handler{services: services}
// }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	expense := router.Group("/expense")
	{
		expense.POST("/", h.createExpense)
	}

 	return router
}
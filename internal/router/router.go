package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/youruser/dexter-transport/internal/app/port"
)

func SetupRouter(r *gin.Engine, h port.Handler) {
	// 1. Init Handlers

	// 2. Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 3. API Routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) { BindReqJson200Resp(c, h.Health(c)) })

		// Task CRUD routes
		v1.POST("/tasks", func(c *gin.Context) { BindReqJson200Resp(c, h.CreateTask(c)) })
		v1.GET("/tasks", func(c *gin.Context) { BindReqJson200Resp(c, h.ListTasks(c)) })
		v1.GET("/tasks/:id", func(c *gin.Context) { BindReqJson200Resp(c, h.GetTask(c)) })
		v1.PUT("/tasks/:id", func(c *gin.Context) { BindReqJson200Resp(c, h.UpdateTask(c)) })
		v1.DELETE("/tasks/:id", func(c *gin.Context) { BindReqJson200Resp(c, h.DeleteTask(c)) })
	}
}

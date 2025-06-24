package interfaces

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, todoHandler *TodoHandler, userHandler *UserHandler) {
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.GET("/:id/todos", userHandler.GetUserWithTodos)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		todos := v1.Group("/todos")
		{
			todos.GET("/:id", todoHandler.GetTodo)
			todos.POST("", todoHandler.CreateTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
			todos.PATCH("/:id/toggle", todoHandler.ToggleTodoStatus)
		}

		userTodos := v1.Group("/user-todos")
		{
			userTodos.GET("/:user_id", todoHandler.GetTodosByUser)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Todo List API is running",
		})
	})
}

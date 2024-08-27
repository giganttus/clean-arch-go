package routes

import (
	"todoapp/handlers"
	"todoapp/repos"
	"todoapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Injection (data flow; Handlers => Services => Repos)
	rep := repos.NewRepos(db)
	svc := services.NewServices(rep)
	hlr := handlers.NewHandlers(svc)

	// Routing group
	apiRouter := router.Group("/api/")

	// Todo routes
	apiRouter.POST("/todo", hlr.CreateTodo)
	apiRouter.GET("/todo", hlr.GetTodos)
	apiRouter.PUT("/todo", hlr.UpdateTodo)
	apiRouter.DELETE("/todo/:id", hlr.DeleteTodo)

	return router
}

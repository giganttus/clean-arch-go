package services

import (
	"todoapp/models"
	"todoapp/repos"

	"github.com/gin-gonic/gin"
)

// Injection of Repository for Service usage
type Services struct {
	Repos repos.IRepos
}

// Constructor (used in router.go)
func NewServices(repos repos.IRepos) *Services {
	return &Services{
		Repos: repos,
	}
}

// Service interface as description, consumed by handlers
type IServices interface {
	// Todos
	CreateTodo(*gin.Context, models.Todos) error
	GetTodos(*gin.Context) ([]models.Todos, error)
	UpdateTodo(*gin.Context, models.Todos) error
	DeleteTodo(*gin.Context, int) error
}

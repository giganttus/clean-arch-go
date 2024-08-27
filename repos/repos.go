package repos

import (
	"todoapp/models"

	"gorm.io/gorm"
)

// Injection of GORM for Repository usage
type Repos struct {
	Db *gorm.DB
}

// Constructor (used in router.go)
func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Db: db,
	}
}

// Repository interface as description, consumed by services
type IRepos interface {
	// Todos - Main functions
	CreateTodo(models.Todos) error
	GetTodos() ([]models.Todos, error)
	UpdateTodo(models.Todos) error
	DeleteTodo(int) error

	// Todos - Support functions
	TodoExists(string) bool
}

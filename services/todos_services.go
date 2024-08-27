package services

import (
	"todoapp/helpers"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func (s *Services) CreateTodo(ctx *gin.Context, input models.Todos) error {
	if todoEx := s.Repos.TodoExists(input.Task); todoEx {
		return helpers.ErrService
	}

	crTodo := models.Todos{
		Task: input.Task,
	}

	return s.Repos.CreateTodo(crTodo)
}

func (s *Services) GetTodos(ctx *gin.Context) ([]models.Todos, error) {
	return s.Repos.GetTodos()
}

func (s *Services) UpdateTodo(ctx *gin.Context, input models.Todos) error {
	if todoEx := s.Repos.TodoExists(input.Task); todoEx {
		return helpers.ErrService
	}

	upTodo := models.Todos{
		ID:        input.ID,
		Task:      input.Task,
		Completed: input.Completed,
	}

	return s.Repos.UpdateTodo(upTodo)
}

func (s *Services) DeleteTodo(ctx *gin.Context, todoID int) error {
	return s.Repos.DeleteTodo(todoID)
}

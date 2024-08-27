package repos

import (
	"todoapp/helpers"
	"todoapp/models"
)

// Main functions
func (r *Repos) CreateTodo(input models.Todos) error {
	if err := r.Db.Create(&input).Error; err != nil {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) GetTodos() ([]models.Todos, error) {
	var todos []models.Todos

	if res := r.Db.Find(&todos); res.Error != nil {
		return nil, helpers.ErrRepo
	}

	return todos, nil
}

func (r *Repos) UpdateTodo(input models.Todos) error {
	res := r.Db.Model(&models.Todos{}).Where("id = ?", input.ID).Updates(&input)
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

func (r *Repos) DeleteTodo(todoID int) error {
	res := r.Db.Where("id = ?", todoID).Delete(&models.Todos{})
	if res.Error != nil || res.RowsAffected == 0 {
		return helpers.ErrRepo
	}

	return nil
}

// Support functions
func (r *Repos) TodoExists(task string) bool {
	res := r.Db.Where("task = ?", task).First(&models.Todos{}).RowsAffected

	return res != 0
}

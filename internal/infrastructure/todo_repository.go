package infrastructure

import (
	"errors"

	"todo-list/internal/domain"
)

type TodoRepositoryImpl struct {
	db *Database
}

func NewTodoRepository(db *Database) domain.TodoRepository {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (r *TodoRepositoryImpl) Create(todo *domain.Todo) error {
	return r.db.DB.Create(todo).Error
}

func (r *TodoRepositoryImpl) GetByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepositoryImpl) GetByUserID(userID uint) ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := r.db.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) Update(todo *domain.Todo) error {
	return r.db.DB.Save(todo).Error
}

func (r *TodoRepositoryImpl) Delete(id uint) error {
	result := r.db.DB.Delete(&domain.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}

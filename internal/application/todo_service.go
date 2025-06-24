package application

import (
	"errors"
	"time"

	"todo-list/internal/domain"
)

type TodoServiceImpl struct {
	repo domain.TodoRepository
}

func NewTodoService(repo domain.TodoRepository) domain.TodoService {
	return &TodoServiceImpl{
		repo: repo,
	}
}

func (s *TodoServiceImpl) CreateTodo(userID uint, title, description string) (*domain.Todo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if userID == 0 {
		return nil, errors.New("user ID cannot be empty")
	}

	todo := &domain.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoServiceImpl) GetTodo(id uint) (*domain.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoServiceImpl) GetTodosByUser(userID uint) ([]domain.Todo, error) {
	todos, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoServiceImpl) UpdateTodo(id uint, title, description string, completed bool) (*domain.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	todo.Completed = completed
	todo.UpdatedAt = time.Now()

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoServiceImpl) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}

func (s *TodoServiceImpl) ToggleTodoStatus(id uint) (*domain.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

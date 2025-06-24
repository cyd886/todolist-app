package application

import (
	"errors"
	"time"

	"todo-list/internal/domain"
)

type UserServiceImpl struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(username, email, password, name string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	if _, err := s.userRepo.GetByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}

	if _, err := s.userRepo.GetByEmail(email); err == nil {
		return nil, errors.New("email already exists")
	}

	user := &domain.User{
		Username:  username,
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) GetUser(id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]domain.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) UpdateUser(id uint, username, email, name string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if username != "" {
		if existingUser, err := s.userRepo.GetByUsername(username); err == nil && existingUser.ID != id {
			return nil, errors.New("username already exists")
		}
		user.Username = username
	}

	if email != "" {
		if existingUser, err := s.userRepo.GetByEmail(email); err == nil && existingUser.ID != id {
			return nil, errors.New("email already exists")
		}
		user.Email = email
	}

	if name != "" {
		user.Name = name
	}

	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserServiceImpl) GetUserWithTodos(id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

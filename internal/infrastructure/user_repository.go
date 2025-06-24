package infrastructure

import (
	"errors"

	"todo-list/internal/domain"
)

type UserRepositoryImpl struct {
	db *Database
}

func NewUserRepository(db *Database) domain.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	return r.db.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.DB.Preload("Todos").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) error {
	return r.db.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	result := r.db.DB.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

package service

import (
	"food-api/api/repository"
	"food-api/models"
)

type FoodService interface {
	GetAll() ([]*models.Food, error)
	Create(food models.Food) (*models.Food, error)
	Delete(id int) (*models.Food, error)
	Update()
	GetById()
}

type FoodServiceImpl struct {
	repo repository.FoodRepository
}

// Create implements FoodService.
func (fs *FoodServiceImpl) Create(food models.Food) (*models.Food, error) {
	result, err := fs.repo.Create(food)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll implements FoodService.
func (fs *FoodServiceImpl) GetAll() ([]*models.Food, error) {
	result, err := fs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Delete implements FoodService.
func (fs *FoodServiceImpl) Delete(id int) (*models.Food, error) {
	result, err := fs.repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements FoodService.
func (fs *FoodServiceImpl) GetById() {
	panic("unimplemented")
}

// Update implements FoodService.
func (fs *FoodServiceImpl) Update() {
	panic("unimplemented")
}

func NewFoodService(repo repository.FoodRepository) FoodService {
	return &FoodServiceImpl{repo: repo}
}

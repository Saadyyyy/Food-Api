package service

import (
	"food-api/api/repository"
	"food-api/models"
)

type FoodService interface {
	GetAll()
	Create(food models.Food) (*models.Food, error)
	Delete()
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

// Delete implements FoodService.
func (fs *FoodServiceImpl) Delete() {
	panic("unimplemented")
}

// GetAll implements FoodService.
func (fs *FoodServiceImpl) GetAll() {
	panic("unimplemented")
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

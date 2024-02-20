package service

import (
	"food-api/api/repository"
	"food-api/models"
)

type FoodService interface {
	GetAll() ([]*models.Food, error)
	Create(food *models.Food) (*models.Food, error)
	Delete(id int) error
	Update(id int, food models.Food) (*models.Food, error)
	GetById(id int) (*models.Food, error)
}

type FoodServiceImpl struct {
	repo repository.FoodRepository
}

// Create implements FoodService.
func (fs *FoodServiceImpl) Create(food *models.Food) (*models.Food, error) {
	// menampung 2 variable dari fungsi repository create
	result, err := fs.repo.Create(food)
	//jika terdapat error akan di handle
	if err != nil {
		return nil, err
	}
	//jika tidak terdapat error akan di return
	return result, nil
}

// GetAll implements FoodService.
func (fs *FoodServiceImpl) GetAll() ([]*models.Food, error) {
	//fungsi get all dari repository menampung 2 variable
	result, err := fs.repo.GetAll()
	//jika fungsi get all dari repository terdapat error akan di handle
	if err != nil {
		return nil, err
	}
	//jika sudah tidak terdapat error akan di return
	return result, nil
}

// Delete implements FoodService.
func (fs *FoodServiceImpl) Delete(id int) error {
	// memanggil fungsi delete dari repository dan di tampung dengan 2 variable
	err := fs.repo.Delete(id)
	//jika saat pemanggilan fungsi terdapat error akan di handle
	if err != nil {
		return err
	}
	//jika sudah tidak terdapat error akan melakukan return
	return nil
}

// GetById implements FoodService.
func (fs *FoodServiceImpl) GetById(id int) (*models.Food, error) {
	//memanggil fungsi get by id dari repository dan menampung 2 variable
	result, err := fs.repo.GetById(id)
	//jika terdapat error akan di handle
	if err != nil {
		return nil, err
	}

	// jika tidak terdapat error akan di return
	return result, nil
}

// Update implements FoodService.
func (fs *FoodServiceImpl) Update(id int, food models.Food) (*models.Food, error) {
	//memanggil fungsi update dari repository dan menampung 2 variable
	data, err := fs.repo.Update(food, id)
	//jika terdapat error kan di handle
	if err != nil {
		return nil, err
	}
	//jika fungsi update tidak terdapat error akan di return
	return data, nil
}

func NewFoodService(repo repository.FoodRepository) FoodService {
	return &FoodServiceImpl{repo: repo}
}

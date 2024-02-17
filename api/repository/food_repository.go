package repository

import (
	"database/sql"
	"fmt"
	"food-api/models"
)

type FoodRepository interface {
	GetAll()
	Create(food *models.Food) (*models.Food, error)
	Delete()
	Update()
	GetById()
}

type FoodRepositoryImpl struct {
	db *sql.DB
}

func NewFoodRepository(db *sql.DB) FoodRepository {
	return &FoodRepositoryImpl{db: db}
}

// Create implements FoodRepository.
func (fr *FoodRepositoryImpl) Create(food *models.Food) (*models.Food, error) {
	_, err := fr.db.Exec("INSERT INTO foods(name,category,price)VALUES($1,$2,$3)", food.Name, food.Category, food.Price)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)

	return food, nil
}

// Delete implements FoodRepository.
func (fr *FoodRepositoryImpl) Delete() {
	panic("unimplemented")
}

// GetAll implements FoodRepository.
func (fr *FoodRepositoryImpl) GetAll() {
	panic("unimplemented")
}

// GetById implements FoodRepository.
func (fr *FoodRepositoryImpl) GetById() {
	panic("unimplemented")
}

// Update implements FoodRepository.
func (fr *FoodRepositoryImpl) Update() {
	panic("unimplemented")
}

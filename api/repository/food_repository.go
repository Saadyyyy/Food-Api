package repository

import (
	"database/sql"
	"food-api/models"
)

type FoodRepository interface {
	GetAll() ([]*models.Food, error)
	Create(food models.Food) (*models.Food, error)
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
func (fr *FoodRepositoryImpl) Create(food models.Food) (*models.Food, error) {
	_, err := fr.db.Exec("INSERT INTO foods(name, category, price) VALUES($1, $2, $3)", food.Name, food.Category, food.Price)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

// GetAll implements FoodRepository.
func (fr *FoodRepositoryImpl) GetAll() ([]*models.Food, error) {
	rows, err := fr.db.Query("SELECT * FROM foods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []*models.Food
	for rows.Next() {
		food := &models.Food{}
		err := rows.Scan(&food.ID, &food.Name, &food.Category, &food.Price)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}

// Delete implements FoodRepository.
func (fr *FoodRepositoryImpl) Delete() {
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

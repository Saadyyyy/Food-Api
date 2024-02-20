package repository

import (
	"database/sql"
	"food-api/models"
)

type FoodRepository interface {
	GetAll() ([]*models.Food, error)
	Create(food *models.Food) (*models.Food, error)
	Delete(id int) error
	Update(food models.Food, id int) (*models.Food, error)
	GetById(id int) (*models.Food, error)
}

type FoodRepositoryImpl struct {
	db *sql.DB
}

// melakukan kontrak dengan

func NewFoodRepository(db *sql.DB) FoodRepository {
	return &FoodRepositoryImpl{db: db}
}

// Create implements FoodRepository.
func (fr *FoodRepositoryImpl) Create(food *models.Food) (*models.Food, error) {
	// query database yang berfuungsi untuk menginputkan table atau row
	// masukkan ke table foods dengan isi(nama,category,price) dengan value(kosong,kosong,kosong << untuk di input di postman) dan mengembalikan ID
	err := fr.db.QueryRow("INSERT INTO foods(name, category, price) VALUES($1, $2, $3) RETURNING id", food.Name, food.Category, food.Price).Scan(&food.ID)
	if err != nil {
		return nil, err
	}
	// jika sudah tidak ada error akan mereturn model.food
	return food, nil
}

// GetAll implements FoodRepository.
func (fr *FoodRepositoryImpl) GetAll() ([]*models.Food, error) {
	//query database untuk mengambil semua isi table database
	//ambil (*)<<semua nilai dari table foods
	rows, err := fr.db.Query("SELECT * FROM foods")
	if err != nil {
		return nil, err
	}
	// fungsi defer akan di jalankan paling bawah
	// fungsi close artinya memasikan jika menutupkan table secara otomati
	defer rows.Close()

	//memanggil model foos
	var foods []*models.Food
	//melakukan perulangan dengan query
	for rows.Next() {
		// memanggil objek model food
		food := &models.Food{}
		// melakukan scan pada nilai nilai dari setiap baris kolom dan menyimpan di objek model food
		err := rows.Scan(&food.ID, &food.Name, &food.Category, &food.Price)
		//jika terdapat err akan di handler
		if err != nil {
			return nil, err
		}
		// memasukkan nilai niali dari objek model food ke dalam slice foods
		foods = append(foods, food)
	}
	// kita melakukan pengecekan apakah ada err selelah memasukkan nilai ke slice
	if err := rows.Err(); err != nil {
		return nil, err
	}
	//jidak sudah tidak ada error maka meretun foods dan err nil
	return foods, nil
}

// Delete implements FoodRepository.
func (fr *FoodRepositoryImpl) Delete(id int) error {
	//query untuk menghapus table berdasarkan id
	//delete dari foods yang dimana id = kosong <<yang dimana akan di input di param postman
	_, err := fr.db.Exec("DELETE FROM foods WHERE id = $1", id)
	//jika query error akan menampilkan err
	if err != nil {
		return err
	}
	//jika tidak ada error akan menampilkan nil
	return nil
}

// GetById implements FoodRepository.
func (fr *FoodRepositoryImpl) GetById(id int) (*models.Food, error) {
	// memanggil objek dari model food
	food := models.Food{}
	//query untuk mengambil row berdasarkan id
	//ambil (*)<<semua nilai dari row dari table foods yang dimana id = kosong <<yang akan di input nanti di param postman
	rows := fr.db.QueryRow("SELECT * FROM foods WHERE id = $1", id)
	//melakukan scan untuk di tampilkan di body json nanti
	err := rows.Scan(&food.ID, &food.Name, &food.Category, &food.Price)
	//jika terdapat err akan di handle
	if err != nil {
		return nil, err
	}
	//jika tidak terdapar error akan mereturn objek dari foods dan err = nil
	return &food, nil
}

// Update implements FoodRepository.
func (fr *FoodRepositoryImpl) Update(food models.Food, id int) (*models.Food, error) {
	//query untuk melakukan delete table foods berdasarkan id
	//update foods di ganti dengan name = kosong, category = kosong, price = kosong yang dimana id = kosong << kosong akan di isi di body postman
	_, err := fr.db.Exec("UPDATE foods SET name = $1, category = $2, price = $3 WHERE id = $4", food.Name, food.Category, food.Price, id)
	//jika terdapat error akan di handle
	if err != nil {
		return nil, err
	}
	//mengantikan id = food id
	//berfungsi untuk saat mengupdate dan menampilkan id sesuai id yang di update
	food.ID = id

	return &food, nil
}

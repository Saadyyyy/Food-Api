package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func InitializeDatabase() (*sql.DB, error) {

	errE := godotenv.Load()
	if errE != nil {
		fmt.Println("Env can connect")
	}

	dbConfig := DatabaseConfig{}
	dbConfig.Host = os.Getenv("DBHOST")
	dbConfig.Port = os.Getenv("DBPORT")
	dbConfig.Username = os.Getenv("DBUSER")
	dbConfig.DBName = os.Getenv("DBNAME")
	dbConfig.Password = os.Getenv("DBPASS")
	// Konfigurasi koneksi database MySQL dengan GORM
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?parseTime=true"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect To Database")
	}

	return db, nil
}

// func Connection() *sql.DB {

// 	// dbConfig := DatabaseConfig{}
// 	// dbConfig.Host = os.Getenv("DBHOST")
// 	// dbConfig.Port = os.Getenv("DBPORT")
// 	// dbConfig.Username = os.Getenv("DBUSER")
// 	// dbConfig.DBName = os.Getenv("DBNAME")
// 	// dbConfig.Password = os.Getenv("DBPASS")
// 	// Konfigurasi koneksi database MySQL dengan GORM
// 	// dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?parseTime=true"

// 	constring := "postgres://postgres:oke@localhost/pqgotest?sslmode=disable"
// 	db, err := sql.Open("postgres", constring)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }

func Connection() (*sql.DB, error) {
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")

	// Buat string koneksi menggunakan nilai dari parameter lingkungan
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Buka koneksi ke database PostgreSQL
	db, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, err
	}

	// Periksa apakah koneksi berhasil
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database")

	return db, nil
}

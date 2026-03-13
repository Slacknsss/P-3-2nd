package config
import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"go-api/models"
)
var DB *gorm.DB
func ConnectDB(){
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
        	os.Getenv("DB_HOST"),
        	os.Getenv("DB_PORT"),
        	os.Getenv("DB_USER"),
        	os.Getenv("DB_NAME"),
	)


db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	panic("Conexion DB echouée :" + err.Error())
}

DB =db
DB.AutoMigrate(&models.User{}, &models.Todo{})
fmt.Println("Connecté a PostgreSQL")
}
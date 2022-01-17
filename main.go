package main

import (
	"cars/cars"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db, err := connectPostgrestDb(false)
	if err != nil {
		log.Fatal(err)
	}
	carsCtrl := cars.CarController{Svc: &cars.CarService{Repo: &cars.CarRepo{DB: db}}}

	r.POST("api/v1/cars", carsCtrl.CreateCar)
	r.GET("api/v1/cars/:id", carsCtrl.GetCarByID)
	r.POST("api/v1/cars/search", carsCtrl.Search)

	err = r.Run() // listen and serve (default port 8080)
	if err != nil {
		log.Fatal(err)
	}
}

// connectPostgrestDb connect to postgres instance
func connectPostgrestDb(sslmode bool) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	s := "disable"
	if sslmode {
		s = "enabled"
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		s,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

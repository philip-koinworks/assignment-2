package routers

import (
	"assignment-2/controllers"
	"assignment-2/models"
	"assignment-2/repositories"
	"assignment-2/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST  = "localhost"
	DB_PORT  = "5432"
	DB_USER  = "postgres"
	DB_PASS  = "postgres"
	DB_NAME  = "assignment2"
	APP_PORT = ":3000"
)

func Router() *gin.Engine {
	db := connectDB()
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	router := gin.Default()

	router.GET("/orders", orderController.GetAllOrders)
	router.POST("/orders", orderController.CreateNewOrder)
	router.DELETE("/order/:id", orderController.DeleteOrder)

	return router
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Default().Println("connection db success")

	db.AutoMigrate(&models.Order{}, &models.Item{})
	return db
}

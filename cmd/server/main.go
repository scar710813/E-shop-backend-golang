package main

import (
	"log"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/configs"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/webserver/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, loadConfigurationError := configs.LoadConfigurations(".")
	if loadConfigurationError != nil {
		log.Fatalf("Error loading configurations: %v", loadConfigurationError)
		panic(loadConfigurationError)
	}

	db, createDatabaseError := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if createDatabaseError != nil {
		log.Fatalf("Error creating database: %v", createDatabaseError)
		panic(createDatabaseError)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDatabase := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDatabase)
	http.HandleFunc("/products", productHandler.CreateProduct)

	println("🔥 Server runing on port 8000")
	http.ListenAndServe(":8000", nil)
}

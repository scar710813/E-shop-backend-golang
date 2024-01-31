package main

import (
	"log"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/configs"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/webserver/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	productDatabase := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDatabase)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products", productHandler.GetAllProducts)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)

	userDatabase := database.NewUser(db)
	userHandler := handler.NewUserHandler(userDatabase)
	router.Post("/users", userHandler.CreateUser)

	println("🔥 Server runing on port 8000")
	http.ListenAndServe(":8000", router)
}

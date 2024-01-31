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
	configuration, loadConfigurationError := configs.LoadConfigurations(".")
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

	router.Route("/products", func(productRouter chi.Router) {
		productDatabase := database.NewProduct(db)
		productHandler := handler.NewProductHandler(productDatabase)

		productRouter.Post("/", productHandler.CreateProduct)
		productRouter.Get("/", productHandler.GetAllProducts)
		productRouter.Get("/{id}", productHandler.GetProduct)
		productRouter.Put("/{id}", productHandler.UpdateProduct)
		productRouter.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Route("/users", func(userRoute chi.Router) {
		userDatabase := database.NewUser(db)
		userHandler := handler.NewUserHandler(
			userDatabase,
			configuration.JwtTokenAuth,
			configuration.JwtExpiresIn,
		)
		userRoute.Post("/", userHandler.CreateUser)
		userRoute.Post("/generate-token", userHandler.GetJwt)
	})

	println("🔥 Server runing on port 8000")
	http.ListenAndServe(":8000", router)
}

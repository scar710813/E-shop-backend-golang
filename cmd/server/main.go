package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PaoloProdossimoLopes/goshop/configs"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/webserver/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(2) * time.Second))

	router.Route("/products", func(productRouter chi.Router) {
		productRouter.Use(jwtauth.Verifier(configuration.JwtTokenAuth))
		productRouter.Use(jwtauth.Authenticator)

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

	const port = 8000
	println(fmt.Sprintf("🔥 Server runing on http://localhost:%v\n\n", port))
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

package main

import (
	"log"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/configs"
	"github.com/PaoloProdossimoLopes/goshop/internal/dto"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/goccy/go-json"
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
	productHandler := NewProductHandler(productDatabase)
	http.HandleFunc("/products", productHandler.CreateProduct)

	println("🔥 Server runing on port 8000")
	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDatabase database.ProductRespository
}

func NewProductHandler(database database.ProductRespository) *ProductHandler {
	return &ProductHandler{
		ProductDatabase: database,
	}
}

func (self *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest dto.CreateProductInput
	jsonDecoderProductError := json.NewDecoder(r.Body).Decode(&productRequest)
	if jsonDecoderProductError != nil {
		http.Error(w, jsonDecoderProductError.Error(), http.StatusBadRequest)
		return
	}

	product, createProductError := entity.NewProduct(productRequest.Name, productRequest.Price)
	if createProductError != nil {
		http.Error(w, createProductError.Error(), http.StatusBadRequest)
		return
	}

	productCreated, createProductError := self.ProductDatabase.Create(product)
	if createProductError != nil {
		http.Error(w, createProductError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productCreated)
}

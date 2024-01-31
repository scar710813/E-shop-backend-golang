package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/internal/dto"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/go-chi/chi"
)

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

func (self *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, getProductError := self.ProductDatabase.FindById(id)
	if getProductError != nil {
		http.Error(w, getProductError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (self *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var productRequest entity.Product
	jsonDecoderProductError := json.NewDecoder(r.Body).Decode(&productRequest)
	if jsonDecoderProductError != nil {
		http.Error(w, jsonDecoderProductError.Error(), http.StatusBadRequest)
		return
	}

	product, getProductError := self.ProductDatabase.FindById(id)
	if getProductError != nil {
		http.Error(w, getProductError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product.Name = productRequest.Name
	product.Price = productRequest.Price

	productUpdated, updateProductError := self.ProductDatabase.Update(product)
	if updateProductError != nil {
		http.Error(w, updateProductError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productUpdated)
}

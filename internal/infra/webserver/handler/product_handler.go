package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/internal/dto"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
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

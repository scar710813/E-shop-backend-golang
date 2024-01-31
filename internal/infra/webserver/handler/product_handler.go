package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// Create product godoc
// @Summary		Create a product
// @Description	Create a product
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.CreateProductInput true "Create product request"
// @Success 	200 {object} entity.Product
// @Failure 	400
// @Failure 	500
// @Router 		/products [post]
// @Security 	ApiKeyAuth
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

// Get product godoc
// @Summary 	Get a product
// @Description Get a product
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Product ID"
// @Success 	200 {object} entity.Product
// @Failure 	400
// @Failure 	500
// @Router 		/products/{id} [get]
// @Security 	ApiKeyAuth
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

// Update product godoc
// @Summary 	Update a product
// @Description Update a product
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Product ID"
// @Param 		request body dto.CreateProductInput true "Update product request"
// @Success 	200 {object} entity.Product
// @Failure 	400
// @Failure 	500
// @Router 		/products/{id} [put]
// @Security 	ApiKeyAuth
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

// Delete product godoc
// @Summary 	Delete a product
// @Description Delete a product
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Product ID"
// @Success 	204
// @Failure 	400
// @Failure 	500
// @Router 		/products/{id} [delete]
// @Security 	ApiKeyAuth
func (self *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	deleteProductError := self.ProductDatabase.Delete(id)
	if deleteProductError != nil {
		http.Error(w, deleteProductError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Get all products godoc
// @Summary 	Get all products
// @Description Get all products
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		page query int false "Page"
// @Param 		limit query int false "Limit"
// @Param 		sort query string false "Sort"
// @Success 	200 {object} []entity.Product
// @Failure 	400
// @Failure 	500
// @Router 		/products [get]
// @Security 	ApiKeyAuth
func (self *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	page, convertPageError := strconv.Atoi(pageParam)
	if convertPageError != nil {
		page = 0
	}

	limit, convertLimitError := strconv.Atoi(limitParam)
	if convertLimitError != nil {
		limit = 10
	}

	products, getProductsError := self.ProductDatabase.FindAll(page, limit, sort)
	if getProductsError != nil {
		http.Error(w, getProductsError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

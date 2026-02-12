package handlers

import (
	"encoding/json"
	"net/http"
	"shop-api/middleware"
	"shop-api/models"
	"shop-api/services"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productService services.ProductService
	shopService    services.ShopService
}

func NewProductHandler(productService services.ProductService, shopService services.ShopService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		shopService:    shopService,
	}
}

// GetAll - GET /products (private - requires auth)
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Get products for the user's shop
	products := h.productService.GetAll(claims.ShopID)

	// Filter response based on role
	if claims.Role == models.RoleSuperAdmin {
		// SuperAdmin sees everything including purchase price
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		// Admin sees everything except purchase price
		var adminProducts []models.AdminProductResponse
		for _, product := range products {
			adminProducts = append(adminProducts, product.ToAdminResponse())
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(adminProducts)
	}
}

type CreateProductRequest struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
	ImageURL      string  `json:"image_url"`
}

// Create - POST /products (private - requires auth)
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Validate
	if req.Name == "" || req.SellingPrice <= 0 {
		http.Error(w, `{"error": "Name and selling price are required"}`, http.StatusBadRequest)
		return
	}

	// Create product with user's ShopID
	product := models.Product{
		Name:          req.Name,
		Description:   req.Description,
		Category:      req.Category,
		PurchasePrice: req.PurchasePrice,
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		ImageURL:      req.ImageURL,
		ShopID:        claims.ShopID,
	}

	created, err := h.productService.Create(product)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return based on role
	if claims.Role == models.RoleSuperAdmin {
		json.NewEncoder(w).Encode(created)
	} else {
		json.NewEncoder(w).Encode(created.ToAdminResponse())
	}
}

// Update - PUT /products/:id (private - requires auth)
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Extract ID from URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		http.Error(w, `{"error": "Invalid URL"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[len(pathParts)-1])
	if err != nil {
		http.Error(w, `{"error": "Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	// Check if product exists and belongs to user's shop
	existing, err := h.productService.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Product not found"}`, http.StatusNotFound)
		return
	}

	if existing.ShopID != claims.ShopID {
		http.Error(w, `{"error": "Unauthorized - product belongs to different shop"}`, http.StatusForbidden)
		return
	}

	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Update product
	product := models.Product{
		Name:          req.Name,
		Description:   req.Description,
		Category:      req.Category,
		PurchasePrice: req.PurchasePrice,
		SellingPrice:  req.SellingPrice,
		Stock:         req.Stock,
		ImageURL:      req.ImageURL,
	}

	updated, err := h.productService.Update(id, product)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Return based on role
	if claims.Role == models.RoleSuperAdmin {
		json.NewEncoder(w).Encode(updated)
	} else {
		json.NewEncoder(w).Encode(updated.ToAdminResponse())
	}
}

// Delete - DELETE /products/:id (private - requires auth)
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Extract ID from URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		http.Error(w, `{"error": "Invalid URL"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[len(pathParts)-1])
	if err != nil {
		http.Error(w, `{"error": "Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	// Check if product exists and belongs to user's shop
	existing, err := h.productService.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Product not found"}`, http.StatusNotFound)
		return
	}

	if existing.ShopID != claims.ShopID {
		http.Error(w, `{"error": "Unauthorized - product belongs to different shop"}`, http.StatusForbidden)
		return
	}

	if err := h.productService.Delete(id); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetPublicProducts - GET /public/:shopID/products (public - no auth required)
func (h *ProductHandler) GetPublicProducts(w http.ResponseWriter, r *http.Request) {
	// Extract shopID from URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		http.Error(w, `{"error": "Invalid URL"}`, http.StatusBadRequest)
		return
	}

	shopID, err := strconv.Atoi(pathParts[1])
	if err != nil {
		http.Error(w, `{"error": "Invalid shop ID"}`, http.StatusBadRequest)
		return
	}

	// Verify shop exists
	shop, err := h.shopService.GetByID(shopID)
	if err != nil {
		http.Error(w, `{"error": "Shop not found"}`, http.StatusNotFound)
		return
	}

	// Get products for this shop
	products := h.productService.GetPublicProducts(shopID)

	// Convert to public response (no purchase price, with WhatsApp link)
	var publicProducts []models.PublicProductResponse
	for _, product := range products {
		// Optionally filter out products with 0 stock
		// if product.Stock > 0 {
		publicProducts = append(publicProducts, product.ToPublicResponse(shop.WhatsAppNumber))
		// }
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publicProducts)
}

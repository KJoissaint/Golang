package handlers

import (
	"encoding/json"
	"net/http"
	"shop-api/middleware"
	"shop-api/models"
	"shop-api/services"
)

type TransactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transactionService services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

type CreateTransactionRequest struct {
	Type      models.TransactionType `json:"type"`
	ProductID *int                   `json:"product_id,omitempty"`
	Quantity  int                    `json:"quantity"`
	Amount    float64                `json:"amount"`
}

// GetAll - GET /transactions (private - requires admin)
func (h *TransactionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	transactions := h.transactionService.GetAll(claims.ShopID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// Create - POST /transactions (private - requires admin)
func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Validate transaction type
	if req.Type != models.TransactionSale &&
		req.Type != models.TransactionExpense &&
		req.Type != models.TransactionWithdrawal {
		http.Error(w, `{"error": "Invalid transaction type"}`, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Quantity <= 0 || req.Amount <= 0 {
		http.Error(w, `{"error": "Quantity and amount must be positive"}`, http.StatusBadRequest)
		return
	}

	// Sales must have a product ID
	if req.Type == models.TransactionSale && req.ProductID == nil {
		http.Error(w, `{"error": "Product ID required for sales"}`, http.StatusBadRequest)
		return
	}

	// Create transaction
	transaction := models.Transaction{
		Type:      req.Type,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Amount:    req.Amount,
		ShopID:    claims.ShopID,
	}

	created, err := h.transactionService.Create(transaction)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// GetDashboard - GET /reports/dashboard (private - SuperAdmin only)
func (h *TransactionHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	stats, err := h.transactionService.GetDashboard(claims.ShopID)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

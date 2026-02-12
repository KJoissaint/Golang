package handlers

import (
	"encoding/json"
	"net/http"
	"shop-api/middleware"
	"shop-api/models"
	"shop-api/services"
)

type ShopHandler struct {
	shopService services.ShopService
}

func NewShopHandler(shopService services.ShopService) *ShopHandler {
	return &ShopHandler{
		shopService: shopService,
	}
}

type UpdateWhatsAppRequest struct {
	WhatsAppNumber string `json:"whatsapp_number"`
}

// UpdateWhatsApp - PUT /shops/whatsapp (SuperAdmin only)
func (h *ShopHandler) UpdateWhatsApp(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Only SuperAdmin can update WhatsApp
	if claims.Role != models.RoleSuperAdmin {
		http.Error(w, `{"error": "Only SuperAdmin can update WhatsApp number"}`, http.StatusForbidden)
		return
	}

	var req UpdateWhatsAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if req.WhatsAppNumber == "" {
		http.Error(w, `{"error": "WhatsApp number is required"}`, http.StatusBadRequest)
		return
	}

	// Update the shop's WhatsApp number
	if err := h.shopService.UpdateWhatsApp(claims.ShopID, req.WhatsAppNumber); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "WhatsApp number updated successfully",
	})
}

// GetAll - GET /shops (private)
func (h *ShopHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	shops := h.shopService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shops)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"shop-api/models"
	"shop-api/services"
)

type AuthHandler struct {
	userService services.UserService
	shopService services.ShopService
}

func NewAuthHandler(userService services.UserService, shopService services.ShopService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		shopService: shopService,
	}
}

type RegisterRequest struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Role     models.Role `json:"role"`
	ShopID   int         `json:"shop_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  models.UserResponse `json:"user"`
	Token string              `json:"token"`
}

// Register - POST /register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Name == "" || req.Email == "" || req.Password == "" {
		http.Error(w, `{"error": "Name, email, and password are required"}`, http.StatusBadRequest)
		return
	}

	// Validate role
	if req.Role != models.RoleSuperAdmin && req.Role != models.RoleAdmin {
		http.Error(w, `{"error": "Invalid role. Must be SuperAdmin or Admin"}`, http.StatusBadRequest)
		return
	}

	// Validate shop exists
	_, err := h.shopService.GetByID(req.ShopID)
	if err != nil {
		http.Error(w, `{"error": "Shop not found"}`, http.StatusBadRequest)
		return
	}

	// Register user
	user, err := h.userService.Register(req.Name, req.Email, req.Password, req.Role, req.ShopID)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user.ToResponse())
}

// Login - POST /login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		http.Error(w, `{"error": "Email and password are required"}`, http.StatusBadRequest)
		return
	}

	// Login
	user, token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	response := AuthResponse{
		User:  user.ToResponse(),
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

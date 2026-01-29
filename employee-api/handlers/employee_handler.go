package handlers

import (
	"employee-api/services"
	"encoding/json"
	"net/http"
)

// gestion des requêtes HTTP liées aux employés
type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

// GetAll 
func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Récupérer les employés depuis le service
	employees := h.service.GetAll()

	// Configurer la réponse en JSON
	w.Header().Set("Content-Type", "application/json")

	// Encoder et envoyer la réponse
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, "Erreur lors de l'encodage JSON", http.StatusInternalServerError)
		return
	}
}

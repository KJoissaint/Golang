package handlers

import (
	"employee-api/models"
	"employee-api/services"
	"encoding/json"
	"net/http"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

// GetAll - GET /employees
func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees := h.service.GetAll()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, "Erreur lors de l'encodage JSON", http.StatusInternalServerError)
		return
	}
}

// Structure pour recevoir les données de création
type CreateEmployeeRequest struct {
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

// Create - POST /employees
func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Lire le body JSON
	var req CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	// Validation minimale
	if req.Name == "" {
		http.Error(w, "Le nom est requis", http.StatusBadRequest)
		return
	}
	if req.Salary < 0 {
		http.Error(w, "Le salaire ne peut pas être négatif", http.StatusBadRequest)
		return
	}

	// Créer l'employé (sans ID, le service le génère)
	employee := models.Employee{
		Name:   req.Name,
		Salary: req.Salary,
	}

	// Appeler le service
	created := h.service.Add(employee)

	// Répondre avec l'employé créé
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// Structure pour recevoir les données d'augmentation
type RaiseSalaryRequest struct {
	ID      int     `json:"id"`
	Percent float64 `json:"percent"`
}

// RaiseSalary - PUT /employees/raise
func (h *EmployeeHandler) RaiseSalary(w http.ResponseWriter, r *http.Request) {
	// Lire le body JSON
	var req RaiseSalaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	// Validation minimale
	if req.ID <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	if req.Percent < 0 {
		http.Error(w, "Le pourcentage ne peut pas être négatif", http.StatusBadRequest)
		return
	}

	// Appeler le service
	if err := h.service.RaiseSalary(req.ID, req.Percent); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Salaire augmenté avec succès"}`))
}

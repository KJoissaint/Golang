package main

import (
	"employee-api/handlers"
	"employee-api/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	employeeService := services.NewEmployeeService()
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	// Route /employees avec dispatch selon la méthode HTTP
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			employeeHandler.GetAll(w, r)
		case http.MethodPost:
			employeeHandler.Create(w, r)
		default:
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	})

	// Route /employees/raise
	http.HandleFunc("/employees/raise", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}
		employeeHandler.RaiseSalary(w, r)
	})

	port := ":8080"
	fmt.Printf("🚀 Serveur démarré sur http://localhost%s\n", port)
	fmt.Println("📍 Endpoints disponibles:")
	fmt.Println("   GET  /employees")
	fmt.Println("   POST /employees")
	fmt.Println("   PUT  /employees/raise")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}

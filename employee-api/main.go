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

	
	http.HandleFunc("/employees", employeeHandler.GetAll)

	// Démarrer le serveur
	port := ":8080"
	fmt.Printf("🚀 Serveur démarré sur http://localhost%s\n", port)
	fmt.Println("📍 Endpoints disponibles:")
	fmt.Println("   GET /employees")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}

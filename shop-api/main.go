package main

import (
	"fmt"
	"log"
	"net/http"
	"shop-api/config"
	"shop-api/handlers"
	"shop-api/middleware"
	"shop-api/services"
	"strings"
)

func main() {
	// Initialize services
	shopService := services.NewShopService()
	userService := services.NewUserService()
	productService := services.NewProductService()
	transactionService := services.NewTransactionService(productService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService, shopService)
	productHandler := handlers.NewProductHandler(productService, shopService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	shopHandler := handlers.NewShopHandler(shopService)

	// Setup routes
	mux := http.NewServeMux()

	// Auth routes (public)
	mux.HandleFunc("/register", methodHandler("POST", authHandler.Register))
	mux.HandleFunc("/login", methodHandler("POST", authHandler.Login))

	// Public routes (no auth required)
	mux.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/products") && r.Method == http.MethodGet {
			productHandler.GetPublicProducts(w, r)
		} else {
			http.Error(w, `{"error": "Not found"}`, http.StatusNotFound)
		}
	})

	// Product routes (private - requires auth)
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.AuthMiddleware(productHandler.GetAll)(w, r)
		case http.MethodPost:
			middleware.AuthMiddleware(productHandler.Create)(w, r)
		default:
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		// Handle /products/:id
		switch r.Method {
		case http.MethodPut:
			middleware.AuthMiddleware(productHandler.Update)(w, r)
		case http.MethodDelete:
			middleware.AuthMiddleware(productHandler.Delete)(w, r)
		default:
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	// Transaction routes (private - requires admin)
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.RequireAdmin(transactionHandler.GetAll)(w, r)
		case http.MethodPost:
			middleware.RequireAdmin(transactionHandler.Create)(w, r)
		default:
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	// Dashboard route (SuperAdmin only)
	mux.HandleFunc(
		"/reports/dashboard",
		methodHandler(
			"GET",
			middleware.RequireSuperAdmin(transactionHandler.GetDashboard),
		),
	)

	// Shop routes
	mux.HandleFunc("/shops", methodHandler("GET",
		middleware.AuthMiddleware(shopHandler.GetAll)))

	mux.HandleFunc("/shops/whatsapp", methodHandler("PUT",
		middleware.RequireSuperAdmin(shopHandler.UpdateWhatsApp)))

	// Root handler - serves static files for non-API routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only serve static files for GET requests
		if r.Method != http.MethodGet {
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		// Serve static files
		http.FileServer(http.Dir("./frontend")).ServeHTTP(w, r)
	})

	// Start server
	fmt.Println("🚀 Shop Management API Server Started")
	fmt.Printf("📍 Server running on http://localhost%s\n", config.ServerPort)
	fmt.Println("\n📋 Available Endpoints:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("\n🔓 PUBLIC ROUTES:")
	fmt.Println("   POST   /register")
	fmt.Println("   POST   /login")
	fmt.Println("   GET    /public/:shopID/products")
	fmt.Println("\n🔒 PRIVATE ROUTES (requires auth):")
	fmt.Println("   GET    /products")
	fmt.Println("   POST   /products")
	fmt.Println("   PUT    /products/:id")
	fmt.Println("   DELETE /products/:id")
	fmt.Println("\n👥 ADMIN ROUTES:")
	fmt.Println("   GET    /transactions")
	fmt.Println("   POST   /transactions")
	fmt.Println("\n👑 SUPER ADMIN ROUTES:")
	fmt.Println("   GET    /reports/dashboard")
	fmt.Println("   PUT    /shops/whatsapp")
	fmt.Println("   GET    /shops")
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("\n📝 Test Accounts:")
	fmt.Println("   SuperAdmin: super@shop1.com / admin123")
	fmt.Println("   Admin:      admin@shop1.com / admin123")
	fmt.Println("\n💡 Tip: Use Authorization header with 'Bearer <token>'")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	if err := http.ListenAndServe(config.ServerPort, corsMiddleware(mux)); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// Helper function to restrict HTTP methods
func methodHandler(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// CORS middleware for development
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

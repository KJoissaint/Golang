package config

import "time"

var (
	// JWT Configuration
	JWTSecret     = []byte("your-secret-key-change-this-in-production")
	JWTExpiration = time.Hour * 24 * 7 // 7 days

	// Server Configuration
	ServerPort = ":8081"
)

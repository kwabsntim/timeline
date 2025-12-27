package main

import (
	"log"
	"net/http"
	"os"
	"timeline/internal/db"
	"timeline/internal/wrap"
)

func main() {
	// 0. Log startup info
	log.Println("🚀 Starting Timeline API...")

	// 1. Initialize database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// 2. Create tables
	err = db.CreateTables(database)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
	log.Println("✅ Database initialized successfully")

	// 3. Initialize repository
	wrapRepo := wrap.NewSqlRepository(database)

	// 4. Initialize service
	wrapService := wrap.NewService(wrapRepo)

	// 5. Initialize handlers
	wrapHandler := wrap.NewHandler(wrapService)

	// 6. Setup routes
	http.HandleFunc("/api/create/wrap", wrapHandler.CreateWrap)
	http.HandleFunc("/api/get/wrap/", wrapHandler.GetWrap)

	// 7. Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 8. Start server
	log.Println("🚀 Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

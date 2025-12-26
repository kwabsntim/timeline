package main

import (
	"log"
	"net/http"
	"timeline/internal/db"
	"timeline/internal/wrap"
)

func main() {
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

	// 7. Start server
	log.Println("🚀 Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

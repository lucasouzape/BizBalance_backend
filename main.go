package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"bizbalance/controllers"
	"bizbalance/repository"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect to PostgreSQL database
	conn, err := repository.ConnectPostgres(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to PostgreSQL database!")

	// Create a controller with the database connection
	itemsController := controllers.NewItemController(conn)

	http.HandleFunc("/pao_de_mel", MethodHandler(itemsController.GetAllPaoDeMel, http.MethodGet))   // Allow GET only
	http.HandleFunc("/pao_de_mel/add", MethodHandler(itemsController.AddPaoDeMel, http.MethodPost)) // Allow POST only

	// Start HTTP server
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// MethodHandler wraps a handler and only allows specified HTTP methods
func MethodHandler(handler http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if r.Method == method {
				handler(w, r)
				return
			}
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

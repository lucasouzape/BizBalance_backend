package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}

/*
package main

import (
    "log"
    "net/http"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/gorilla/mux"
    "path/to/yourproject/controllers"
)

func main() {
    db := InitDB()
    router := mux.NewRouter()

    // Rota de exemplo
    router.HandleFunc("/items", controllers.GetItems(db)).Methods("GET")
    log.Println("API is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func InitDB() *gorm.DB {
    dsn := "host=localhost user=user password=password dbname=mydatabase port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    return db
}

*/

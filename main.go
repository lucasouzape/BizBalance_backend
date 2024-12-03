package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"bizbalance/controllers"
	"bizbalance/repository"
)

type ProductData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r)
	}
}

func getProductData(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		products, err := repository.FetchProductData(conn)
		if err != nil {
			http.Error(w, "Erro ao buscar dados dos produtos", http.StatusInternalServerError)
			return
		}

		labels := []string{}
		values := []float64{}
		for _, product := range products {
			labels = append(labels, product.Name)
			values = append(values, product.Value)
		}

		data := ProductData{
			Labels: labels,
			Values: values,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando valores padrão.")
	}

	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			log.Fatalf("Variável de ambiente obrigatória ausente: %s", key)
		}
	}
}

func main() {
	loadEnv()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	conn, err := repository.ConnectPostgres(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer conn.Close()

	itemsController := controllers.NewItemController(conn)

	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "./frontend"
	}
	fmt.Printf("Servindo arquivos estáticos de: %s\n", frontendPath)
	fs := http.FileServer(http.Dir(frontendPath))
	http.Handle("/", fs)

	apiBasePath := "/api"
	http.HandleFunc(apiBasePath+"/pao_de_mel", enableCORS(itemsController.GetAllPaoDeMel))
	http.HandleFunc(apiBasePath+"/pao_de_mel/add", enableCORS(itemsController.AddPaoDeMel))
	http.HandleFunc(apiBasePath+"/products-data", enableCORS(getProductData(conn)))

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

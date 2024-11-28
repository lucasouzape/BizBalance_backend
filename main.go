package main

import (
	"encoding/json"
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

// Middleware para habilitar CORS
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

// Carregar variáveis de ambiente com valores padrão
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando valores padrão.")
	}

	// Verificar se variáveis essenciais estão definidas
	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			log.Fatalf("Variável de ambiente obrigatória ausente: %s", key)
		}
	}
}

func main() {
	// Carregar variáveis de ambiente
	loadEnv()

	// Obter credenciais do banco de dados
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Conectar ao banco de dados
	conn, err := repository.ConnectPostgres(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer conn.Close()
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Criar controlador
	itemsController := controllers.NewItemController(conn)

	// Servir arquivos estáticos do frontend
	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "./frontend"
	}
	fs := http.FileServer(http.Dir(frontendPath))
	http.Handle("/", fs)

	// Rotas da API
	apiBasePath := "/api"
	http.HandleFunc(apiBasePath+"/pao_de_mel", enableCORS(itemsController.GetAllPaoDeMel))
	http.HandleFunc(apiBasePath+"/pao_de_mel/add", enableCORS(itemsController.AddPaoDeMel))
	http.HandleFunc(apiBasePath+"/calculate", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		// Verificar o método HTTP
		if r.Method == http.MethodGet {
			// Responder com uma mensagem explicativa para GET
			response := map[string]string{
				"message": "Use o método POST para enviar os dados para o cálculo.",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			// Encaminhar para o controlador Calculate para POST
			itemsController.Calculate(w, r)
		}
	}))

	// Iniciar servidor HTTP
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// PaoDeMel representa a estrutura de um registro de pao_de_mel
type PaoDeMel struct {
	ID              int       `json:"id"`
	Sabor           string    `json:"sabor"`
	Quantidade      int       `json:"quantidade"`
	Validade        time.Time `json:"validade"`
	PrecoCusto      float64   `json:"preco_custo"`
	PrecoVenda      float64   `json:"preco_venda"`
	AdministradorID int       `json:"administrador_id"`
}

// ItemController gerencia requisições HTTP.
type ItemController struct {
	DB *sql.DB
}

// NewItemController cria uma nova instância de ItemController com conexão ao banco de dados.
func NewItemController(db *sql.DB) *ItemController {
	return &ItemController{DB: db}
}

// GetAllPaoDeMel recupera todos os itens do banco de dados.
func (c *ItemController) GetAllPaoDeMel(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, sabor, quantidade, validade, preco_custo, preco_venda, administrador_id FROM pao_de_mel")
	if err != nil {
		http.Error(w, "Erro ao buscar dados: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao buscar dados: %v", err)
		return
	}
	defer rows.Close()

	var items []PaoDeMel
	for rows.Next() {
		var item PaoDeMel
		if err := rows.Scan(&item.ID, &item.Sabor, &item.Quantidade, &item.Validade, &item.PrecoCusto, &item.PrecoVenda, &item.AdministradorID); err != nil {
			http.Error(w, "Erro ao processar linha: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao processar linha: %v", err)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// AddPaoDeMel adiciona um novo item ao banco de dados.
func (c *ItemController) AddPaoDeMel(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao ler o corpo da requisição: %v", err)
		return
	}
	defer r.Body.Close()

	var item PaoDeMel
	if err := json.Unmarshal(body, &item); err != nil {
		http.Error(w, "Formato JSON inválido: "+err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar JSON: %v", err)
		return
	}

	// Validação dos dados recebidos
	if item.Quantidade <= 0 {
		http.Error(w, "Quantidade deve ser maior que zero", http.StatusBadRequest)
		log.Printf("Erro: Quantidade inválida %d", item.Quantidade)
		return
	}
	if item.PrecoCusto <= 0 {
		http.Error(w, "Preço de custo deve ser maior que zero", http.StatusBadRequest)
		log.Printf("Erro: Preço de custo inválido %.2f", item.PrecoCusto)
		return
	}

	// Definir valores padrão para campos ausentes
	item.AdministradorID = 1 // ID padrão do administrador

	if item.Validade.IsZero() {
		item.Validade = time.Now().AddDate(0, 0, 30) // Validade padrão de 30 dias
		log.Printf("Validade calculada: %v", item.Validade)
	}

	if item.PrecoVenda == 0 {
		item.PrecoVenda = item.PrecoCusto * 1.4 // Margem de 40% se não fornecido
		log.Printf("PrecoVenda calculado: %.2f", item.PrecoVenda)
	}

	// Inserir novo registro no banco de dados
	query := `
		INSERT INTO pao_de_mel (sabor, quantidade, validade, preco_custo, preco_venda, administrador_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err = c.DB.QueryRow(query, item.Sabor, item.Quantidade, item.Validade, item.PrecoCusto, item.PrecoVenda, item.AdministradorID).Scan(&item.ID)
	if err != nil {
		http.Error(w, "Falha ao inserir dados: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao inserir no banco de dados: %v", err)
		return
	}

	// Log do item inserido com sucesso
	log.Printf("Item inserido com sucesso: ID=%d, Validade=%v, PrecoVenda=%.2f", item.ID, item.Validade, item.PrecoVenda)

	// Responder com o item criado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Calculate lida com o cálculo solicitado pelo frontend.
func (c *ItemController) Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Responder com uma mensagem explicativa para o método GET
		response := map[string]string{
			"message": "Use o método POST para enviar os dados para o cálculo.",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Tratar método POST
	var requestData struct {
		QuantidadeVendida int     `json:"quantidade_vendida"`
		PrecoCusto        float64 `json:"preco_custo"`
		PrecoVenda        float64 `json:"preco_venda"`
	}

	// Analisar o corpo da solicitação
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Formato JSON inválido: "+err.Error(), http.StatusBadRequest)
		log.Printf("Erro ao decodificar JSON no cálculo: %v", err)
		return
	}

	// Validação do cálculo
	if requestData.QuantidadeVendida <= 0 || requestData.PrecoVenda <= requestData.PrecoCusto {
		http.Error(w, "Dados inválidos para cálculo", http.StatusBadRequest)
		log.Printf("Erro nos dados do cálculo: Quantidade=%d, PreçoCusto=%.2f, PreçoVenda=%.2f",
			requestData.QuantidadeVendida, requestData.PrecoCusto, requestData.PrecoVenda)
		return
	}

	// Realizar o cálculo
	retorno := (requestData.PrecoVenda - requestData.PrecoCusto) * float64(requestData.QuantidadeVendida)

	// Responder com o resultado
	response := map[string]interface{}{
		"retorno": retorno,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

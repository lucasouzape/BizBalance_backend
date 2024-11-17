package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// PaoDeMel represents the structure of a pao_de_mel record
type PaoDeMel struct {
	ID              int       `json:"id"`
	Sabor           string    `json:"sabor"`
	Quantidade      int       `json:"quantidade"`
	Validade        time.Time `json:"validade"`
	PrecoCusto      float64   `json:"preco_custo"`
	PrecoVenda      float64   `json:"preco_venda"`
	AdministradorID int       `json:"administrador_id"`
}

// MyController handles HTTP requests.
type ItemController struct {
	DB *sql.DB
}

// NewItemController creates a new instance of ItemController with a database connection.
func NewItemController(db *sql.DB) *ItemController {
	return &ItemController{DB: db}
}

// GetItems is an example HTTP handler.
func (c *ItemController) GetAllPaoDeMel(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, sabor, quantidade, validade, preco_custo, preco_venda, administrador_id FROM pao_de_mel")
	if err != nil {
		http.Error(w, "Failed to fetch data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []PaoDeMel
	for rows.Next() {
		var item PaoDeMel
		if err := rows.Scan(&item.ID, &item.Sabor, &item.Quantidade, &item.Validade, &item.PrecoCusto, &item.PrecoVenda, &item.AdministradorID); err != nil {
			http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetItems is an example HTTP handler.
func (c *ItemController) AddPaoDeMel(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var item PaoDeMel
	if err := json.Unmarshal(body, &item); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set Administrator ID to 1
	item.AdministradorID = 1

	// Insert the new record into the database
	query := `
		INSERT INTO pao_de_mel (sabor, quantidade, validade, preco_custo, preco_venda, administrador_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err = c.DB.QueryRow(query, item.Sabor, item.Quantidade, item.Validade, item.PrecoCusto, item.PrecoVenda, item.AdministradorID).Scan(&item.ID)
	if err != nil {
		http.Error(w, "Failed to insert data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

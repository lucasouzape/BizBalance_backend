package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Product representa a estrutura dos dados do produto
type Product struct {
	Name  string
	Value float64
}

// ConnectPostgres establishes a connection to the PostgreSQL database.
func ConnectPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the database to ensure connection is established
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}

// FetchProductData retrieves product data from the database
func FetchProductData(conn *sql.DB) ([]Product, error) {
	query := "SELECT name, value FROM products" // Ajuste para o nome correto da tabela no banco de dados
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Value); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		products = append(products, product)
	}

	// Verificar se houve algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return products, nil
}

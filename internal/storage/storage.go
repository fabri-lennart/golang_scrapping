package storage

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

// Product represents a single scraped product entity
type Product struct {
	ProductName string
	Price       string
}

// Database holds the active SQL connection pool
type Database struct {
	Conn *sql.DB
}

// NewDatabase initializes a new SQLite connection and creates the required tables
func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Verify the connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT NOT NULL,
		price TEXT NOT NULL,
		scraped_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	// 1. AQUÍ CERRAMOS NewDatabase devolviendo la base de datos creada
	return &Database{Conn: db}, nil
} // <-- Esta llave cierra NewDatabase

// 2. AQUÍ DECLARAMOS SaveProduct, totalmente independiente y usando el tipo "Product"
func (db *Database) SaveProduct(p Product) error {
	query := `INSERT INTO products (product_name, price) VALUES (?, ?)`

	// db.Conn.Exec executes the query without returning any rows
	_, err := db.Conn.Exec(query, p.ProductName, p.Price)
	return err
}

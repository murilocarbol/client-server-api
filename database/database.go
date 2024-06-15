package database

import (
	"context"
	"database/sql"
	"time"

	// driver for SQLite3 Database
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	dbName        = "cotacoes.db"
	dbSaveTimeout = 10 * time.Millisecond
)

func UpDatabase() {

	// Estabelecer a conexão com o banco de dados SQLite3
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar se a conexão com o banco de dados é válida
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// criacao da tabela cotacao
	query := `CREATE TABLE IF NOT EXISTS cotacao (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		timestamp TEXT
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("erro ao criar tabela: %v", err)
	}
}

func SaveCotacao(bid string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), dbSaveTimeout)
	defer cancel()

	query := "INSERT INTO cotacao (bid, timestamp) VALUES (?, ?)"
	_, err = db.ExecContext(ctx, query, bid, time.Now().Format(time.RFC822))
	if err != nil {
		return err
	}
	return nil
}

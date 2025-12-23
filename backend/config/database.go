package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL_CONECTION")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	DB = conn

	var now time.Time
	err = DB.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
	createTable()
}

func createTable() {
	ctx := context.Background()
	query := `
    CREATE TABLE IF NOT EXISTS items (
        id SERIAL PRIMARY KEY,
        ticker VARCHAR(10) NOT NULL,
        target_from VARCHAR(10),
        target_to VARCHAR(10),
        company TEXT,
        action TEXT,
        brokerage TEXT,
        rating_from TEXT,
        rating_to TEXT,
        time TEXT,
        created_at TIMESTAMP DEFAULT current_timestamp()
    )`

	_, err := DB.Exec(ctx, query)
	if err != nil {
		log.Fatal("Error creando tabla:", err)
	}
	log.Println("Tabla products lista")
}

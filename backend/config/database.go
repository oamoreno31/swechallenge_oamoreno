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
	seedTestData()
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

func seedTestData() {
	ctx := context.Background()

	// Check if data already exists
	var count int
	err := DB.QueryRow(ctx, "SELECT COUNT(*) FROM items").Scan(&count)
	if err != nil {
		log.Printf("Error checking existing data: %v", err)
		return
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Printf("Database already contains %d items. Skipping seed data.", count)
		return
	}

	// Insert test data
	testData := []struct {
		ticker      string
		target_from string
		target_to   string
		company     string
		action      string
		brokerage   string
		rating_from string
		rating_to   string
		time        string
	}{
		{
			ticker:      "AAPL",
			target_from: "150",
			target_to:   "180",
			company:     "Apple Inc.",
			action:      "Buy",
			brokerage:   "Goldman Sachs",
			rating_from: "Hold",
			rating_to:   "Buy",
			time:        time.Now().Format(time.RFC3339),
		},
		{
			ticker:      "MSFT",
			target_from: "300",
			target_to:   "350",
			company:     "Microsoft Corporation",
			action:      "Buy",
			brokerage:   "Morgan Stanley",
			rating_from: "Hold",
			rating_to:   "Buy",
			time:        time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "GOOGL",
			target_from: "120",
			target_to:   "140",
			company:     "Alphabet Inc.",
			action:      "Hold",
			brokerage:   "JPMorgan Chase",
			rating_from: "Buy",
			rating_to:   "Hold",
			time:        time.Now().Add(-48 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "AMZN",
			target_from: "140",
			target_to:   "160",
			company:     "Amazon.com Inc.",
			action:      "Buy",
			brokerage:   "Bank of America",
			rating_from: "Hold",
			rating_to:   "Buy",
			time:        time.Now().Add(-72 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "TSLA",
			target_from: "200",
			target_to:   "250",
			company:     "Tesla, Inc.",
			action:      "Strong Buy",
			brokerage:   "Wells Fargo",
			rating_from: "Buy",
			rating_to:   "Strong Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		}, {
			ticker:      "RMTI",
			target_from: "$3.00",
			target_to:   "$2.50",
			company:     "Rockwell Medical",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "AKBA",
			target_from: "$8.00",
			target_to:   "$6.00",
			company:     "Akebia Therapeutics",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "CECO",
			target_from: "$52.00",
			target_to:   "$57.00",
			company:     "CECO Environmental",
			action:      "target raised by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "BLND",
			target_from: "$3.50",
			target_to:   "$3.00",
			company:     "Blend Labs",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Neutral",
			rating_to:   "Neutral",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "VYGR",
			target_from: "$30.00",
			target_to:   "$25.00",
			company:     "Voyager Therapeutics",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "BCBP",
			target_from: "$9.50",
			target_to:   "$9.00",
			company:     "BCB Bancorp, Inc. (NJ)",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Market Perform",
			rating_to:   "Market Perform",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "DEFT",
			target_from: "$8.00",
			target_to:   "$3.00",
			company:     "DeFi Technologies",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "LAMR",
			target_from: "$135.00",
			target_to:   "$145.00",
			company:     "Lamar Advertising",
			action:      "target raised by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "AMP",
			target_from: "$568.00",
			target_to:   "$554.00",
			company:     "Ameriprise Financial",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Buy",
			rating_to:   "Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ticker:      "MLSS",
			target_from: "$1.25",
			target_to:   "$1.00",
			company:     "Milestone Scientific",
			action:      "target lowered by",
			brokerage:   "NONE",
			rating_from: "Speculative Buy",
			rating_to:   "Speculative Buy",
			time:        time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
	}

	query := `
		INSERT INTO items (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	inserted := 0
	for _, data := range testData {
		_, err := DB.Exec(ctx, query,
			data.ticker,
			data.target_from,
			data.target_to,
			data.company,
			data.action,
			data.brokerage,
			data.rating_from,
			data.rating_to,
			data.time,
		)
		if err != nil {
			log.Printf("Error inserting test data for %s: %v", data.ticker, err)
			continue
		}
		inserted++
	}

	if inserted > 0 {
		log.Printf("Successfully inserted %d test items into the database", inserted)
	}
}

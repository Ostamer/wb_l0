package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS orders (
		order_uid TEXT PRIMARY KEY,
		data JSONB NOT NULL
	)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Connected to PostgreSQL and ensured table exists")
}

func saveOrderToDB(order Order) {
	data, err := json.Marshal(order)
	if err != nil {
		log.Println("Failed to marshal order:", err)
		return
	}

	_, err = db.Exec("INSERT INTO orders (order_uid, data) VALUES ($1, $2) ON CONFLICT (order_uid) DO NOTHING",
		order.OrderUID, data)
	if err != nil {
		log.Println("Failed to insert order into DB:", err)
	}
}

func loadCacheFromDB() {
	rows, err := db.Query("SELECT data FROM orders")
	if err != nil {
		log.Println("Failed to load cache from DB:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			log.Println("Failed to read row:", err)
			continue
		}

		var order Order
		if err := json.Unmarshal(data, &order); err != nil {
			log.Println("Failed to unmarshal order:", err)
			continue
		}

		cacheLock.Lock()
		cache[order.OrderUID] = order
		cacheLock.Unlock()
	}

	log.Println("Cache loaded from DB")
}

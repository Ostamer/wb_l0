package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func startKafkaConsumer() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     kafkaTopic,
		GroupID:   "order-service",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer r.Close()

	log.Println("Kafka consumer started and subscribed to topic:", kafkaTopic)

	for {

		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			continue
		}

		log.Printf("Message received: %s\n", string(m.Value))

		var order Order
		err = json.Unmarshal(m.Value, &order)
		if err != nil {
			log.Printf("Invalid message format: %v", err)
			continue
		}

		saveOrderToDB(order)

		cacheLock.Lock()
		cache[order.OrderUID] = order
		cacheLock.Unlock()

		log.Printf("Order %s processed and cached\n", order.OrderUID)
	}
}

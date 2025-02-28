package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	// Get Redis address from environment variables (or default to localhost)
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Check connectivity to Redis
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// HTTP handler to increment and display visitor count
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment the visitor counter in Redis
	count, err := client.Incr("visitors").Result()
	if err != nil {
		http.Error(w, "Error incrementing visitors", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Visitor Count: %d", count)
}

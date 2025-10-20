package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			return
		}
		
		next(w, r)
	}
}

func main() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASS", "postgres"),
		getEnv("DB_NAME", "demo"),
	)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	http.HandleFunc("/api/add", corsMiddleware(addItem))
	http.HandleFunc("/api/delete", corsMiddleware(deleteItem))
	http.HandleFunc("/api/list", corsMiddleware(listItems))
	http.HandleFunc("/api/load", corsMiddleware(simulateLoad))
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	port := getEnv("PORT", "8080")
	log.Printf("Backend running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func addItem(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing ?name=", 400)
		return
	}
	_, err := db.Exec("INSERT INTO items (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Added %s\n", name)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing ?id=", 400)
		return
	}
	_, err := db.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Deleted id=%s\n", id)
}

func listItems(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, name FROM items ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var it Item
		rows.Scan(&it.ID, &it.Name)
		items = append(items, it)
	}
	json.NewEncoder(w).Encode(items)
}

// func simulateLoad(w http.ResponseWriter, _ *http.Request) {
// 	start := time.Now()
	
// 	size := 250

// 	for k := 0; k < 100; k++ {
// 		matrixA := make([][]float64, size)
// 		matrixB := make([][]float64, size)
// 		result := make([][]float64, size)
		
// 		for i := 0; i < size; i++ {
// 			matrixA[i] = make([]float64, size)
// 			matrixB[i] = make([]float64, size)
// 			result[i] = make([]float64, size)
// 			for j := 0; j < size; j++ {
// 				matrixA[i][j] = rand.Float64()
// 				matrixB[i][j] = rand.Float64()
// 			}
// 		}
		
// 		for i := 0; i < size; i++ {
// 			for j := 0; j < size; j++ {
// 				for m := 0; m < size; m++ {
// 					result[i][j] += matrixA[i][m] * matrixB[m][j]
// 				}
// 			}
// 		}
// 	}
	
// 	duration := time.Since(start)
// 	fmt.Fprintf(w, "Load simulated in %v\n", duration)
// }

func simulateLoad(w http.ResponseWriter, _ *http.Request) {
	start := time.Now()
	
	numWorkers := runtime.NumCPU() // Используем все доступные ядра
	var wg sync.WaitGroup
	
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Каждый воркер выполняет тяжелые вычисления
			for i := 0; i < 10_000_000; i++ {
				val := rand.Float64()
				_ = math.Sin(val) * math.Cos(val) + math.Tan(val) * math.Log(val+1)
				_ = math.Sin(val) * math.Cos(val) + math.Tan(val) * math.Log(val+1)
				_ = math.Sqrt(val * val + float64(i)) * math.Exp(val)
			}
		}(w)
	}
	
	wg.Wait()
	duration := time.Since(start)
	fmt.Fprintf(w, "Parallel CPU load (%d workers) completed in %v\n", numWorkers, duration)
}

func getEnv(k, def string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return def
}

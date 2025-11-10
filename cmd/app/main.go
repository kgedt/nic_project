// package main

// import (
// 	"io"
// 	"log"
// 	"net/http"
// )

// func healthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	io.WriteString(w, "ok")
// }

// func main() {

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/healthz", healthCheck)

// 	srv := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}

// 	log.Println("server started on :8080")
// 	if err := srv.ListenAndServe(); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// 1) грузим .env (если есть)
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty (проверь .env или переменные окружения)")
	}

	// 2) создаём пул с таймаутом на подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("pgxpool.New: %v", err)
	}
	defer pool.Close()

	// 3) проверяем, что коннект жив (Ping)
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	// 4) пробный запрос
	var now time.Time
	if err := pool.QueryRow(ctx, "SELECT NOW()").Scan(&now); err != nil {
		log.Fatalf("QueryRow NOW() failed: %v", err)
	}

	fmt.Println("✅ Connected! Server time:", now.Format(time.RFC3339))
}

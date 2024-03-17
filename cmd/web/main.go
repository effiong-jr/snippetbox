package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/effiong-jr/snippetbox/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	logInfo  *log.Logger
	logError *log.Logger
	snippets *models.SnippetModel
}

func main() {

	dbConnectionString := "postgres://web:54321@localhost:5432/snippetbox"

	dbPool, err := pgxpool.New(context.Background(), dbConnectionString)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Unable to create connection pool: %v", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	logError := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		logInfo:  logInfo,
		logError: logError,
		snippets: &models.SnippetModel{DB: dbPool},
	}

	addr := flag.String("addr", ":4000", "HTTP server address")

	flag.Parse()

	// mux := http.NewServeMux()

	mux := app.routes()

	logInfo.Printf("Server started on port %s", *addr)

	err = http.ListenAndServe(*addr, mux)

	logError.Fatal(err)

}

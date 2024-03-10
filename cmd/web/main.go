package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	logInfo  *log.Logger
	logError *log.Logger
}

func main() {

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	logError := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		logInfo:  logInfo,
		logError: logError,
	}

	addr := flag.String("addr", ":4000", "HTTP server address")

	flag.Parse()

	// mux := http.NewServeMux()

	mux := app.routes()

	logInfo.Printf("Server started on port %s", *addr)

	err := http.ListenAndServe(*addr, mux)

	logError.Fatal(err)

}

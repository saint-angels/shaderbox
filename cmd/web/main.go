package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type contextKey string
const contextKeyIsAuthenticated = contextKey("isAuthenticated")


type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)



	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,

	}

	infoLog.Printf("starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
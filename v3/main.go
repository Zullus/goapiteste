package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var stdLog, errLog *log.Logger

func main() {

	stdLog = log.New(os.Stdout, "INFO ", 0)
	stdLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	errLog = log.New(os.Stderr, "ERROR ", 1)
	errLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/erro", errorSampleHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		panic(err)

	}
}

func errorSampleHandler(w http.ResponseWriter, r *http.Request) {
	errLog.Fatal("Deu erro")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Olá Mundo</h1><div>Bem vindo!</div>")
}

func logRequest(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		stdLog.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		handler.ServeHTTP(w, r)

	})
}

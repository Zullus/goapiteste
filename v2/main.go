package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		panic(err)

	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Olá Mundo</h1><div>Bem vindo!</div>")
}

func logRequest(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		handler.ServeHTTP(w, r)

	})
}

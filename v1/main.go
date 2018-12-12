package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		panic(err)

	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Olá Mundo</h1><div>Bem vindo!</div>")
}

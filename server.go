package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("El servidor esta a la escucha del puerto 5000")
	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, "Pagina no encontrada")
	} else {
		template.Execute(w, nil)
	}
}

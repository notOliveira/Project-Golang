package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Função genérica para renderizar templates
func renderTemplate(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		http.Error(w, "Erro ao carregar página", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Resgatando query params
		queryParamTest := r.URL.Query().Get("paramTest")

		if queryParamTest == "" {
			queryParamTest = "No paramTest"
		}
		fmt.Printf("%s\n\n\n", queryParamTest)
		renderTemplate(w, "index.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html")
	})

	// Servindo arquivos estáticos (CSS, JS, imagens)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Inicializando o servidor
	fmt.Println("Server is running on port 8080.\nPress Ctrl + C to stop the server.")
	http.ListenAndServe(":8000", nil)
}

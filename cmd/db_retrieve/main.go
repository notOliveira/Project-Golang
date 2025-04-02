package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	
	"projectGO/pkg/database"
)

type Usuario struct {
	ID    int
	Nome  string
	Idade int
}

func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nome, idade FROM usuarios")
	if err != nil {
		http.Error(w, "Erro ao buscar os dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Idade); err != nil {
			http.Error(w, "Erro ao ler os dados", http.StatusInternalServerError)
			return
		}
		usuarios = append(usuarios, u)
	}

	tmpl, err := template.ParseFiles("cmd/db_retrieve/templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, usuarios)
}

func adicionarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		idade, _ := strconv.Atoi(r.FormValue("idade"))

		_, err := database.DB.Exec("INSERT INTO usuarios (nome, idade) VALUES (?, ?)", nome, idade)
		if err != nil {
			http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func excluirUsuario(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := database.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Erro ao excluir usuário", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	database.InitDB() // Inicializa a conexão com o banco

	http.HandleFunc("/", listarUsuarios)
	http.HandleFunc("/adicionar", adicionarUsuario)
	http.HandleFunc("/excluir", excluirUsuario)

	log.Println("Servidor rodando em :8000")
	http.ListenAndServe(":8000", nil)
}

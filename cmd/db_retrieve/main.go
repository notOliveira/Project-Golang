package main

import (
	//"fmt"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projectGO/pkg/database"
)

type Usuario struct {
	ID    int
	Nome  string
	Idade int
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/db_retrieve/templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil) // sem passar os dados aqui
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

	// Define o header como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func adicionarUsuario(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    defer r.Body.Close()

    var usuario Usuario

    // Tenta decodificar o JSON da requisição
    err := json.NewDecoder(r.Body).Decode(&usuario)

	log.Println("Dados recebidos:", usuario)
    if err != nil {
        log.Println("Erro ao decodificar JSON:", err)
        http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
        return
    }

    // Log para ver os dados recebidos
    log.Println("Nome:", usuario.Nome)
    log.Println("Idade:", usuario.Idade)

    // Insere os dados no banco
    _, err = database.DB.Exec("INSERT INTO Usuarios (nome, idade) VALUES (?, ?)", usuario.Nome, usuario.Idade)
    if err != nil {
        log.Println("Erro ao inserir usuário:", err)
        http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
        return
    }

    // Redireciona após sucesso
    http.Redirect(w, r, "/", http.StatusSeeOther)
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

	http.HandleFunc("/", home)
	http.HandleFunc("/listar", listarUsuarios)
	http.HandleFunc("/adicionar", adicionarUsuario)
	http.HandleFunc("/excluir", excluirUsuario)

	log.Println("Servidor rodando em :8000")
	http.ListenAndServe(":8000", nil)
}

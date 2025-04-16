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

// Função para renderizar templates HTML
func renderTemplate(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		http.Error(w, "Erro ao carregar página", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Endpoint GET de usuários
func getUsers(w http.ResponseWriter, r *http.Request) {
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

// Endpoint POST para adicionar um usuário
func addUser(w http.ResponseWriter, r *http.Request) {
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

	// Insere os dados no banco
	_, err = database.DB.Exec("INSERT INTO Usuarios (nome, idade) VALUES (?, ?)", usuario.Nome, usuario.Idade)
	if err != nil {
		log.Println("Erro ao inserir usuário:", err)
		http.Error(w, "Erro ao inserir usuário", http.StatusInternalServerError)
		return
	}

	// Retorna um JSON do usuário inserido
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// Endpoint DELETE para excluir um usuário
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	// Extrai o ID do usuário da URL
	id := r.URL.Path[len("/deleteUser/"):]
	if id == "" {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Executa a exclusão no banco de dados
	_, err := database.DB.Exec("DELETE FROM Usuarios WHERE id = ?", id)
	if err != nil {
		log.Println("Erro ao excluir usuário:", err)
		http.Error(w, "Erro ao excluir usuário", http.StatusInternalServerError)
		return
	}
}

func main() {
	database.InitDB() // Inicializa a conexão com o banco

	// Páginas HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})

	http.HandleFunc("/adicionar", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "adicionar.html")
	})

	// Endpoints API
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/addUser", addUser)
	http.HandleFunc("/deleteUser/{id}", deleteUser)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Servidor rodando em :8000")
	http.ListenAndServe(":8000", nil)
}

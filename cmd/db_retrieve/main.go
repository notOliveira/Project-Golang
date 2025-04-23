package main

import (
	//"fmt"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projectGO/pkg/database"
	"strconv"
)

type Usuario struct {
	ID    int
	Nome  string
	Idade int
}

func jsonError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
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

// Endpoint PUT para atualizar um usuário
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		jsonError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	idStr := r.URL.Path[len("/updateUser/"):]
	if idStr == "" {
		jsonError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "ID deve ser um número inteiro")
		return
	}

	var usuario Usuario
	err = json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		jsonError(w, http.StatusBadRequest, "Erro ao decodificar JSON")
		return
	}

	log.Println("Dados recebidos:", usuario)

	// Verifica se o usuário existe
	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Usuarios WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		log.Println("Erro ao verificar existência do usuário:", err)
		jsonError(w, http.StatusInternalServerError, "Erro interno ao verificar usuário")
		return
	}
	if !exists {
		jsonError(w, http.StatusNotFound, "Usuário não encontrado")
		return
	}

	// Atualiza os dados no banco
	_, err = database.DB.Exec("UPDATE Usuarios SET nome = ?, idade = ? WHERE id = ?", usuario.Nome, usuario.Idade, id)
	if err != nil {
		log.Println("Erro ao atualizar usuário:", err)
		jsonError(w, http.StatusInternalServerError, "Erro ao atualizar usuário")
		return
	}

	usuario.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
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

	http.HandleFunc("/atualizar", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "atualizar.html")
	})

	// Endpoints API
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/addUser", addUser)
	http.HandleFunc("/deleteUser/{id}", deleteUser)
	http.HandleFunc("/updateUser/{id}", updateUser)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Servidor rodando em :8000")
	http.ListenAndServe(":8000", nil)
}

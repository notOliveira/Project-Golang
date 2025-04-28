package main

import (
	//"fmt"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projectGO/pkg/database"
	"strconv"
	"database/sql"
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

// GET /getUsers
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nome, idade FROM usuarios")
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao buscar os dados")
		return
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Idade); err != nil {
			jsonError(w, http.StatusInternalServerError, "Erro ao ler os dados")
			return
		}
		usuarios = append(usuarios, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// GET /getUser/{id}
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	idStr := r.URL.Path[len("/getUser/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		jsonError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var usuario Usuario
	err = database.DB.QueryRow("SELECT id, nome, idade FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nome, &usuario.Idade)
	if err != nil {
		if err == sql.ErrNoRows {
			jsonError(w, http.StatusNotFound, "Usuário não encontrado")
		} else {
			jsonError(w, http.StatusInternalServerError, "Erro ao buscar usuário")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// POST /addUser
func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}
	defer r.Body.Close()

	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil || usuario.Nome == "" || usuario.Idade <= 0 {
		jsonError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	res, err := database.DB.Exec("INSERT INTO Usuarios (nome, idade) VALUES (?, ?)", usuario.Nome, usuario.Idade)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao inserir usuário")
		return
	}

	lastID, _ := res.LastInsertId()
	usuario.ID = int(lastID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// DELETE /deleteUser/{id}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		jsonError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	idStr := r.URL.Path[len("/deleteUser/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		jsonError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Verifica se o usuário existe
	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM usuarios WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao verificar usuário")
		return
	}
	if !exists {
		jsonError(w, http.StatusNotFound, "Usuário não encontrado")
		return
	}

	_, err = database.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao excluir usuário")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário excluído com sucesso"})
}

// PUT /updateUser/{id}
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		jsonError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	idStr := r.URL.Path[len("/updateUser/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		jsonError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var usuario Usuario
	err = json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil || usuario.Nome == "" || usuario.Idade <= 0 {
		jsonError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM usuarios WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao verificar usuário")
		return
	}
	if !exists {
		jsonError(w, http.StatusNotFound, "Usuário não encontrado")
		return
	}

	_, err = database.DB.Exec("UPDATE usuarios SET nome = ?, idade = ? WHERE id = ?", usuario.Nome, usuario.Idade, id)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Erro ao atualizar usuário")
		return
	}

	usuario.ID = id
	w.Header().Set("Content-Type", "application/json")
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

	http.HandleFunc("/atualizar/{id}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "atualizar.html")
	})

	// Endpoints API
	http.HandleFunc("/getUsers", getUsers)
	http.HandleFunc("/getUser/{id}", getUser)
	http.HandleFunc("/addUser", addUser)
	http.HandleFunc("/deleteUser/{id}", deleteUser)
	http.HandleFunc("/updateUser/{id}", updateUser)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Servidor rodando em :8000")
	http.ListenAndServe(":8000", nil)
}

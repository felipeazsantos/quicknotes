package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while executing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/noteView.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while executing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Criando uma nova nota...")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()
	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}

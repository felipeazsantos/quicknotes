package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
		"views/templates/pages/note-view.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while executing template: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func noteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Criando a nota")
}

func main() {
	config := loadConfig()

	slog.SetDefault(newLogger(os.Stdout, config.GetLevelLog()))

	slog.Info(fmt.Sprintf("Servidor rodando na porta %s", config.ServerPort))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux))
}

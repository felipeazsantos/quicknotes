package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/felipeazsantos/quicknotes/internal/handlers/errorapp"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return errors.New("not found")
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return fmt.Errorf("error while parsing template: %s", err.Error())
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		return fmt.Errorf("error while executing template: %s", err.Error())
	}

	return nil
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return errorapp.WithStatus(http.StatusBadRequest, "note not found")
	}

	if id == "0" {
		return errorapp.WithStatus(http.StatusNotFound, "note not found")
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return errorapp.WithStatus(http.StatusInternalServerError, "error while parsing template: %s", err.Error())
	}

	err = t.ExecuteTemplate(w, "base", id)
	if err != nil {
		return errorapp.WithStatus(http.StatusInternalServerError, "error while executing template: %s", err.Error())
	}

	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return fmt.Errorf("error while parsing template: %s", err.Error())
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		return fmt.Errorf("error while executing template: %s", err.Error())
	}

	return nil
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return fmt.Errorf("method not allowed: %s", r.Method)
	}

	fmt.Fprint(w, "Criando a nota")

	return nil
}

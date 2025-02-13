package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/felipeazsantos/quicknotes/internal/handlers/errorapp"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var statusErr errorapp.StatusError
		if errors.As(err, &statusErr) {
			if statusErr.StatusCode() == http.StatusNotFound {
				files := []string{
					"views/templates/base.html",
					"views/templates/pages/404.html",
				}

				t, err := template.ParseFiles(files...)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = t.ExecuteTemplate(w, "base", nil)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			http.Error(w, statusErr.Error(), statusErr.StatusCode())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

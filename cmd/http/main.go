package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipeazsantos/quicknotes/internal/handlers"
)

func main() {
	config := loadConfig()

	slog.SetDefault(newLogger(os.Stdout, config.GetLevelLog()))
	slog.Info(fmt.Sprintf("Servidor rodando na porta %s", config.ServerPort))

	noteHandler := handlers.NewNoteHandler()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	mux.Handle("/", handlers.HandlerWithError(noteHandler.NoteList))
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.Handle("/note/new", handlers.HandlerWithError(noteHandler.NoteNew))
	mux.Handle("/note/create", handlers.HandlerWithError(noteHandler.NoteCreate))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux))
}

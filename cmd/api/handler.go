package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	data := map[string]string{
		"status":  "available",
		"version": version,
		"env":     app.config.env,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err2 := w.Write(js)
	if err2 != nil {
		http.Error(w, "internal error: write json", http.StatusInternalServerError)
		return
	}
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.getBook(w, r)
	} else if r.Method == http.MethodPost {
		// todo
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.getBook(w, r)
	} else if r.Method == http.MethodPost {
		app.updateBook(w, r)
	} else if r.Method == http.MethodDelete {
		app.deleteBook(w, r)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookIdFromRequest(r)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	books, err := app.models.Books.GetById(id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	jsonError := app.writeJson(books, w)
	if jsonError != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookIdFromRequest(r)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Update book with ID %d", id)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookIdFromRequest(r)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete book with ID %d", id)
}

func parseBookIdFromRequest(r *http.Request) (int64, error) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %v", err)
	}
	return idInt, nil
}

func (app *application) writeJson(data any, w http.ResponseWriter) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err2 := w.Write(js)
	if err2 != nil {
		return err2
	}

	return nil
}

// Package controllers handles HTTP request routing and response handling for the bookstore API.
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/PremModhaOfficial/go-bookstore/pkg/models"
	"github.com/PremModhaOfficial/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	bookDetails, err := models.GetBookByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var fillerBook models.Book

	utils.ParseBody(r, &fillerBook)

	b := fillerBook.CreateBook()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	deletedBook, err := models.DeleteBook(bookID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var newBook models.Book

	vars := mux.Vars(r)
	bookIDStr := vars["bookId"]
	utils.ParseBody(r, &newBook)
	bookID, err := strconv.ParseInt(bookIDStr, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	oldBook, err := models.GetBookByID(bookID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if newBook.Name != "" {
		oldBook.Name = newBook.Name
	}
	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}
	if newBook.Publication != "" {
		oldBook.Publication = newBook.Publication
	}
	oldBook.UpdatedAt = time.Now()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(oldBook)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewBooks)
}

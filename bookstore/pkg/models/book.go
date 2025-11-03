package models

import (
	"errors"
	"sync"
	"time"
)

type Book struct {
	ID          int64      `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name"`
	Author      string     `json:"author"`
	Publication string     `json:"publication"`
}

var (
	books  = make(map[int64]*Book)
	mu     sync.RWMutex
	nextID int64 = 1
)

func (b *Book) CreateBook() *Book {
	mu.Lock()
	defer mu.Unlock()
	now := time.Now()
	b.ID = nextID
	nextID++
	b.CreatedAt = now
	b.UpdatedAt = now
	books[b.ID] = b
	return b
}

func GetAllBooks() []Book {
	mu.RLock()
	defer mu.RUnlock()
	var result []Book
	for _, book := range books {
		if book.DeletedAt == nil {
			result = append(result, *book)
		}
	}
	return result
}

func GetBookById(Id int64) (*Book, error) {
	mu.RLock()
	defer mu.RUnlock()
	book, exists := books[Id]
	if !exists || book.DeletedAt != nil {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func DeleteBook(Id int64) (*Book, error) {
	mu.Lock()
	defer mu.Unlock()
	book, exists := books[Id]
	if !exists || book.DeletedAt != nil {
		return nil, errors.New("book not found")
	}
	now := time.Now()
	book.DeletedAt = &now
	return book, nil
}

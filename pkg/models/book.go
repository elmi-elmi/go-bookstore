package models

import (
	"github.com/elmi-elmi/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Author      string `json:"author" gorm:"column:author"`
	Publication string `json:"publication" gorm:"column:publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Create a new book
func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}

// Get all books
func GetAllBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// Get a single book by ID
func GetBookById(ID int64) (*Book, *gorm.DB, error) {
	var book Book
	result := db.First(&book, ID) // cleaner than Where+Find
	if result.Error != nil {
		return nil, nil, result.Error
	}
	return &book, db, nil
}

// Delete a book by ID
func DeleteBook(ID int64) error {
	result := db.Delete(&Book{}, ID) // no need to fetch first
	return result.Error
}

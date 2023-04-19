package models

import (
	"github.com/Nurmanbetov/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book model for creating book units.
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author`
	Publication string `json:"publication"`
}

// Function that connects to database.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Book method for creation of the Book units.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Function which return all book items from database.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Function returns Book item by ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// Function to delete Book item by ID from database.
func DeleteBook(ID int64) Book {
	var book Book

	db.Where("ID=?", ID).Delete(book)
	return book
}

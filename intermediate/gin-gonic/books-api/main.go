package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin" // Ignore this wraning as we use module for path
)

// Define Our Struct for Book Library
type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Qty    uint   `json:qty`
}

var books = []Book{
	{ID: "1", Name: "Golang for Everyone", Author: "Ahmed", Qty: 100},
	{ID: "2", Name: "Gin-Gonic for Creating Api", Author: "Sharique", Qty: 10},
	{ID: "3", Name: "World Going to Doom ", Author: "CPN", Qty: 40},
	{ID: "4", Name: "Love Your own self", Author: "Chirag", Qty: 110},
	{ID: "5", Name: "Self claim opportunity", Author: "Ahmed", Qty: 50},
}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(id string) (*Book, error) {
	for key, value := range books {
		if value.ID == id {
			return &books[key], nil
		}
	}
	return nil, errors.New("Book doesn't exist")
}
func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messege": "Book doesn't exist"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messege": "Book doesn't add"})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messege": "Query for Book checkout doesn't exist"})
		return
	}
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messege": " Book doesn't exist"})
		return
	}
	if book.Qty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messege": " Book Not available"})
		return
	}
	book.Qty -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messege": "Query for Book return doesn't exist"})
		return
	}
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messege": " Book doesn't exist"})
		return
	}
	if book.Qty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messege": " Book Not Found"})
		return
	}
	book.Qty += 1
	c.IndentedJSON(http.StatusOK, book)
}
func main() {
	router := gin.Default()
	// Simple group: v1
	routerV1 := router.Group("/api/v1")
	{
		routerV1.GET("/books/list", getAllBooks)
		routerV1.GET("/books/:id", getBookById)
		routerV1.POST("/books/add", AddBook)
		routerV1.PATCH("/books/checkout", checkoutBook)
		routerV1.PATCH("/books/return", returnBook)

	}
	router.Run("localhost:8080")
}

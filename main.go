package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"` // Serialize to JSON
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	// Capitalize the first letter of each word to make it seen by other modules and be readable as Json
}

var books = []Book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 1},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 2},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 3},
}

func main() {
	route := gin.Default()                 // Responsible for handling different routes and endpoint apis
	route.GET("/books", getBooks)          // Get request to get books
	route.POST("/books", createBooks)      // Post request to append book (will be cached only in the ram for now)
	route.GET("/books:id", getBookByID)    // Get request to get a book by id
	route.PATCH("/books:id", checkOutBook) // Patch request to check out the book if it is available
	route.Run("localhost:8080")            // Router will listen and serve on port 8080
}

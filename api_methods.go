package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	// gin.Context is a struct that contains information about the request
	c.JSON(http.StatusOK, books)
	// IndentedJSON with first parameter is the status code and second is the data to be sent, also StatusOK is the default status code with 200 coming from http module
}

func createBooks(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Message back on failure with that body
		return
	}
	books = append(books, newBook)                   // Append the new book to the books slice
	c.JSON(http.StatusOK, gin.H{"success": newBook}) // Message back on success with that body
}

func getBookByID(c *gin.Context) {
	id := c.Param("id") // Get the id from the url
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func checkOutBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			if book.Quantity > 0 { // if the book is available
				book.Quantity--
				books[i] = book
				c.JSON(http.StatusOK, gin.H{"success": book})
				return
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not available"}) // If the book is not available
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": "Book deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

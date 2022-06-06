# **Go API**

<p align="center">
    <img src="https://miro.medium.com/max/1246/1*hV308VnNWS1xlrSaztOHkw.png" alt="GO" />
</p>

---

## How To Use ⬇️

Use the command below to launch the server in the terminal

```bash
go run main.go api_methods.go
```
# Routes:
```go
    route.GET("/books", getBooks)                    // Get request to get books
	route.POST("/books", createBooks)                // Post request to append book (will be cached only in the ram for now)
	route.GET("/books/:id", getBookByID)             // Get request to get a book by id
	route.PATCH("/books/checkout/:id", checkOutBook) // Patch request to check out the book if it is available
	route.DELETE("/books/:id", deleteBook)           // Delete request to delete a book
```
---

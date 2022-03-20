package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AndreasPB/go-library-api/library"
)

func main() {
	router := gin.Default()
	router.GET("/books", library.GetBooks)
	router.GET("/books/:id", library.BookById)
	router.POST("/books", library.CreateBook)
	router.PATCH("/checkout", library.CheckoutBook)
	router.PATCH("/return", library.ReturnBook)
	router.Run("localhost:8080")
}

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"github.com/AndreasPB/go-library-api/library"
)

func main() {
	router := gin.Default()

	// Allows everything
	router.Use(cors.Default())

	router.GET("/books", library.GetBooks)
	router.GET("/books/:id", library.BookById)
	router.POST("/books", library.CreateBook)
	router.PATCH("/checkout", library.CheckoutBook)
	router.PATCH("/return", library.ReturnBook)

	router.Run("0.0.0.0:8080")
}

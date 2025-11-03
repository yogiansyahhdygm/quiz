package main

import (
	"quiz/controllers"
	"quiz/database"
	"quiz/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Perpustakaan API")
	})

	api := r.Group("/api")

	// semua CRUD buku & kategori butuh Basic Auth
	api.Use(middlewares.BasicAuthMiddleware())
	{
		// Buku
		api.GET("/books", controllers.GetAllBuku)
		api.GET("/books/:id", controllers.GetBukuByID)
		api.POST("/books", controllers.CreateBuku)
		api.PUT("/books/:id", controllers.UpdateBuku)
		api.DELETE("/books/:id", controllers.DeleteBuku)

		// Kategori
		api.GET("/categories", controllers.GetAllKategori)
		api.GET("/categories/:id", controllers.GetKategoriByID)
		api.GET("/categories/:id/books", controllers.GetBukuByKategori)
		api.POST("/categories", controllers.CreateKategori)
		api.PUT("/categories/:id", controllers.UpdateKategori)
		api.DELETE("/categories/:id", controllers.DeleteKategori)

	}

	r.Run(":8080")
}
